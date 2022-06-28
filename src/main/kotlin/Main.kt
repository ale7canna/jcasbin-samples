
import models.CustomSubject
import org.casbin.jcasbin.main.CoreEnforcer.newModel
import org.casbin.jcasbin.main.Enforcer
import org.casbin.jcasbin.persist.file_adapter.FileAdapter
import kotlin.system.measureTimeMillis


fun main(args: Array<String>) {
    println("Hello World!")

    // Try adding program arguments via Run/Debug configuration.
    // Learn more about running applications: https://www.jetbrains.com/help/idea/running-applications.html.
    println("Program arguments: ${args.joinToString()}")

    val modelText = {}.javaClass.getResource("model.conf")?.readText()
    val adapterStream = {}.javaClass.getResource("policy.csv")?.openStream()
    val model = newModel(modelText)
    val adapter = FileAdapter(adapterStream)

    val enforcer = Enforcer(model, adapter, true)

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

            val sub = CustomSubject(name=name, isAdmin=isAdmin) // the user that wants to access a resource.
            println(enforcer.enforce(sub, domain, obj, act))
        }
    }
    println("Evaluation of ${nPolicies} policies took ${timeInMillis} ms")
}