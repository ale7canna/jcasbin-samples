import models.CustomSubject
import org.casbin.adapter.JDBCAdapter
import org.casbin.jcasbin.main.CoreEnforcer.newModel
import org.casbin.jcasbin.main.Enforcer
import kotlin.system.measureTimeMillis


fun main(args: Array<String>) {
    println("Hello World!")

    // Try adding program arguments via Run/Debug configuration.
    // Learn more about running applications: https://www.jetbrains.com/help/idea/running-applications.html.
    println("Program arguments: ${args.joinToString()}")

    val dbDriver = "org.postgresql.Driver"
    val dbUrl = "jdbc:postgresql://localhost:6543/jcasbin-sample"
    val dbUsername = "postgres"
    val dbPassword = "postgres"

    val modelText = {}.javaClass.getResource("model.conf")?.readText()
    val model = newModel(modelText)
    val adapter = JDBCAdapter(dbDriver, dbUrl, dbUsername, dbPassword);

    val enforcer = Enforcer(model, adapter, true)
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

    println("Enforce correctly set up!")

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
    println("Evaluation of ${nPolicies} policies took ${timeInMillis} ms")
}