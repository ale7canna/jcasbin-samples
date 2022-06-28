package commands

import models.CustomSubject
import utils.getEnforcer
import kotlin.system.measureTimeMillis

fun dbSetup(args: List<String>) {
    val enforcer = getEnforcer()
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

fun benchmark(args: List<String>) {
    val enforcer = getEnforcer()

    val names = listOf("ale", "alice", "bob")
    val actions = listOf("read", "write")
    val objects = listOf("course1", "course2", "data1", "data2", "course", "exam", "exam1", "exam2", "content")
    val domains = listOf("domain1", "domain2")
    val nPolicies = 100
    val timeInMillis = measureTimeMillis {
        for (k in 0..nPolicies) {
            val name = names.shuffled().first()
            val domain = domains.shuffled().first()
            val obj = objects.shuffled().first()
            val act = actions.shuffled().first()
            val isAdmin = listOf(true, false).shuffled().first()

            val sub = CustomSubject(name = name, isAdmin = isAdmin) // the user that wants to access a resource.
            println(enforcer.enforce(sub, domain, obj, act))
        }
    }
    println("Evaluation of $nPolicies policies took $timeInMillis ms")
}

fun checkPolicy(args: List<String>) {
    if (args.size < 4) {
        println("not enough arguments $args")
        return
    }

    val enforcer = getEnforcer()
    val subject = args[0]
    val domain = args[1]
    val obj = args[2]
    val action = args[3]
    val sub = CustomSubject(name = subject, isAdmin = false) // the user that wants to access a resource.
    println("Checking policy $args. Result: ${enforcer.enforce(sub, domain, obj, action)}")
}