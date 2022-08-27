package benchmarks

import commands.getDomains
import commands.getNames
import commands.getResources
import models.CustomSubject
import org.casbin.jcasbin.main.Enforcer
import org.openjdk.jmh.annotations.*
import utils.getEnforcer
import java.util.concurrent.*

@State(Scope.Benchmark)
@Fork(1)
@Measurement(iterations = 3, time = 5, timeUnit = TimeUnit.SECONDS)
class CasbinBenchmark {
    private lateinit var enforcer: Enforcer

    val names = getNames()
    val actions = listOf("read", "edit", "consume", "share")
    val objects = getResources()
    val domains = getDomains()

//    @Warmup(iterations = 1, timeUnit = TimeUnit.SECONDS)
//    @Setup
    fun setUp() {
        enforcer = getEnforcer(null, null)
    }

//    @Benchmark
    fun checkPolicy(): Boolean {
        val name = names.shuffled().first()
        val domain = domains.shuffled().first()
        val obj = objects.shuffled().first()
        val act = actions.shuffled().first()
        val isAdmin = listOf(true, false).shuffled().first()

        val sub = CustomSubject(name = name, isAdmin = isAdmin) // the user that wants to access a resource.
        return enforcer.enforce(sub, domain, obj, act)
    }
}