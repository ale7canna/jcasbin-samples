
import org.casbin.jcasbin.main.CoreEnforcer.newModel
import org.casbin.jcasbin.main.Enforcer
import org.casbin.jcasbin.persist.file_adapter.FileAdapter


fun main(args: Array<String>) {
    println("Hello World!")

    // Try adding program arguments via Run/Debug configuration.
    // Learn more about running applications: https://www.jetbrains.com/help/idea/running-applications.html.
    println("Program arguments: ${args.joinToString()}")

    val modelText = {}.javaClass.getResource("model.conf")?.readText()
    val adapterStream = {}.javaClass.getResource("policy.csv")?.openStream()
    val model = newModel(modelText)
    val adapter = FileAdapter(adapterStream)

    val enforcer = Enforcer(model, adapter)

    val sub = "alice" // the user that wants to access a resource.
    val domain = "domain1"
    val obj = "data1" // the resource that is going to be accessed.
    val act = "read" // the operation that the user performs on the resource.

    println(enforcer.enforce(sub, domain, obj, act))
}