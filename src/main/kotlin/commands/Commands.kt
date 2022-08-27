package commands

import models.CustomSubject
import org.casbin.adapter.JDBCAdapter
import org.casbin.jcasbin.main.Enforcer
import org.casbin.jcasbin.persist.file_adapter.FilteredAdapter
import utils.getEnforcer
import utils.getModel
import utils.getPostgresAdapter
import java.util.*
import kotlin.system.measureTimeMillis

fun dbSetup(args: List<String>) {
    val enforcer = getEnforcer(null, null)
    enforcer.addNamedGroupingPolicies(
        "g", listOf(
            listOf("alice", "admin", "domain1"),
            listOf("bob", "admin", "domain2"),
            listOf("ale", "admin", "domain1"),
        )
    )
    enforcer.addNamedGroupingPolicies(
        "g2", listOf(
            listOf("content", "root"),
            listOf("course", "content"),
            listOf("exam", "content"),
            listOf("course1", "course"),
            listOf("course2", "course"),
            listOf("exam1", "exam"),
            listOf("exam2", "exam")
        )
    )
    enforcer.addPolicies(
        listOf(
            listOf("admin", "*", "data1", "read"),
            listOf("admin", "*", "data1", "write"),
            listOf("admin", "domain2", "data2", "read"),
            listOf("admin", "domain2", "data2", "write"),
            listOf("admin", "*", "content", "*"),
        )
    )
    enforcer.savePolicy()
    println("db setup completed")
}

fun dbSetupLarge(args: List<String>) {
    val enforcer = getEnforcer(null, null)
    val names = getNames()
    val roles = getRoles()
    val domains = getDomains()
    val pattern = """(?i)\d""".toRegex()
    val resources = getResources().map {listOf(it, it.take(pattern.find(it)?.range?.first ?: it.length))}
    val userRoles =
        generateSequence {
            listOf(
                names.shuffled().first(),
                roles.shuffled().first(),
                domains.shuffled().first()
            )
        }.take(1000).toList()
    val types = listOf("course", "exam", "quiz", "lab")
    val genType = {
        val t = types.shuffled().first()
        listOf(t + UUID.randomUUID(), t)
    }
    val randomResources = generateSequence { genType() }.take(10000).toList()
    val resourceRoles = listOf(
        listOf("content", "root"),
        listOf("course", "content"),
        listOf("exam", "content"),
        listOf("lab", "content"),
    )
    val actions = listOf("read", "edit", "consume", "share")
    val policies = generateSequence {
        listOf(
            (names.plus("manager")).shuffled().first(),
            domains.shuffled().first(),
            resources.shuffled().first()[0],
            actions.shuffled().first()
        )
    }.take(1000).toList()
    enforcer.addNamedGroupingPolicies(
        "g", userRoles
    )
    enforcer.addNamedGroupingPolicies(
        "g2", resourceRoles + resources + randomResources
    )
    enforcer.addPolicies(policies)
    enforcer.savePolicy()
    println("db setup completed")
}

fun benchmark(args: List<String>) {
    val start = System.currentTimeMillis()

//    val enforcer = getEnforcer(null, null)
    println("enforcer init took ${System.currentTimeMillis() - start} ms")
    val model = getModel()
    val adapter = getPostgresAdapter() as JDBCAdapter

    val names = getNames()
    val actions = listOf("read", "edit", "consume", "share")
    val objects = getResources()
    val domains = getDomains()
    val nPolicies = if (args.size == 1) args[0].toInt() else 100
    val timeInMillis = measureTimeMillis {
        for (k in 0..nPolicies) {
            val name = names.shuffled().first()
            val domain = domains.shuffled().first()
            val obj = objects.shuffled().first()
            val act = actions.shuffled().first()
            val isAdmin = listOf(true, false).shuffled().first()

            val sub = CustomSubject(name = name, isAdmin = isAdmin) // the user that wants to access a resource.

            val policyFilter = FilteredAdapter.Filter()
            policyFilter.g = arrayOf("", "")
            policyFilter.p = arrayOf(
                name, "", "", ""
            )
            val enforcer = Enforcer(model)
            enforcer.adapter = adapter
            enforcer.loadFilteredPolicy(policyFilter)
            val result = enforcer.enforce(sub, domain, obj, act)
//            println(result)
        }
    }
    println("Evaluation of $nPolicies policies took $timeInMillis ms")
}

fun checkPolicy(args: List<String>) {
    if (args.size < 4) {
        println("not enough arguments $args")
        return
    }

    val enforcer = getEnforcer(null, null)
    val subject = args[0]
    val domain = args[1]
    val obj = args[2]
    val action = args[3]
    val sub = CustomSubject(name = subject, isAdmin = false) // the user that wants to access a resource.
    println("Checking policy $args. Result: ${enforcer.enforce(sub, domain, obj, action)}")
}