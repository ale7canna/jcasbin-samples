import commands.benchmark
import commands.checkPolicy
import commands.dbSetup


val commands = hashMapOf<String, (List<String>) -> Unit>(
    "db-setup" to ::dbSetup,
    "benchmark" to ::benchmark,
    "check" to ::checkPolicy,
)

fun main(args: Array<String>) {
    println("Hello World!")
    println("Program arguments: ${args.joinToString()}")

    var input = readln().split(" ")
    while (input[0] != "exit") {
        val command = input.firstOrNull()
        if (command != null && commands.containsKey(command))
            commands[command]?.let { it(input.drop(1)) }
        else
            println("No command found $command")
        input = readln().split(" ")
    }
}

