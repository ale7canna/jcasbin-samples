import commands.benchmark
import commands.checkPolicy
import commands.dbSetup


val commands = hashMapOf<String, (List<String>) -> Unit>(
    "db-setup" to ::dbSetup,
    "benchmark" to ::benchmark,
    "check" to ::checkPolicy,
)

fun main(args: Array<String>) {
    println("Program arguments: ${args.joinToString()}")

    if (args.isNotEmpty()) {
        runCommand(args.toList())
        return
    }
    var input = readln().split(" ")
    while (input[0] != "exit") {
        runCommand(input)
        input = readln().split(" ")
    }
}

private fun runCommand(input: List<String>) {
    val command = input.firstOrNull()
    if (command != null && commands.containsKey(command))
        commands[command]?.let { it(input.drop(1)) }
    else
        println("No command found $command")
}

