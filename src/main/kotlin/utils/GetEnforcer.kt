package utils

import org.casbin.adapter.JDBCAdapter
import org.casbin.jcasbin.main.CoreEnforcer.newModel
import org.casbin.jcasbin.main.Enforcer

fun getEnforcer(): Enforcer {
    val dbHost = System.getenv("DB_HOST") ?: "localhost"
    val dbPort = System.getenv("DB_PORT") ?: "6543"
    val dbDriver = "org.postgresql.Driver"
    val dbUrl = "jdbc:postgresql://$dbHost:$dbPort/jcasbin-sample"
    val dbUsername = "postgres"
    val dbPassword = "postgres"

    val modelText = object {}.javaClass.getResource("/model.conf")?.readText()
    val model = newModel(modelText)
    val adapter = JDBCAdapter(dbDriver, dbUrl, dbUsername, dbPassword);

    return Enforcer(model, adapter, true)
}