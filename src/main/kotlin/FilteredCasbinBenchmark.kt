package benchmarks

import commands.getDomains
import commands.getNames
import commands.getResources
import models.CustomSubject
import org.casbin.adapter.JDBCAdapter
import org.casbin.jcasbin.main.Enforcer
import org.casbin.jcasbin.persist.file_adapter.FilteredAdapter
import org.openjdk.jmh.annotations.*
import utils.getEnforcer
import utils.getModel
import utils.getPostgresAdapter
import java.util.concurrent.*

@State(Scope.Benchmark)
@Fork(1)
@Warmup(iterations = 0)
@Measurement(iterations = 2, time = 10, timeUnit = TimeUnit.SECONDS)
class FilteredCasbinBenchmark {
    val names = getNames()
    val actions = listOf("read", "edit", "consume", "share")
    val objects = getResources()
    val domains = getDomains()

    @Setup
    @Warmup(iterations = 1)
    fun setUp() {
        Thread.sleep(1000)
    }

    @Benchmark
    fun checkPolicy(): Boolean {
        val name = names.shuffled().first()
        val domain = domains.shuffled().first()
        val obj = objects.shuffled().first()
        val act = actions.shuffled().first()
        val isAdmin = listOf(true, false).shuffled().first()

        val sub = CustomSubject(name = name, isAdmin = isAdmin) // the user that wants to access a resource.

        val model = getModel()
        val adapter = getPostgresAdapter() as JDBCAdapter
        val policyFilter = FilteredAdapter.Filter()
        policyFilter.g = arrayOf("", "")
        policyFilter.p = arrayOf(
            name, "", "", ""
        )
        adapter.loadFilteredPolicy(model, policyFilter)
        val enforcer = Enforcer(model, adapter)
//        enforcer.adapter = adapter
//        enforcer.loadFilteredPolicy(policyFilter)

        return enforcer.enforce(sub, domain, obj, act)
    }
}