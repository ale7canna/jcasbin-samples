package utils

import com.googlecode.aviator.Main
import org.casbin.adapter.JDBCAdapter
import org.casbin.jcasbin.main.CoreEnforcer.newModel
import org.casbin.jcasbin.main.Enforcer

fun getEnforcer(): Enforcer {
    val dbDriver = "org.postgresql.Driver"
    val dbUrl = "jdbc:postgresql://localhost:6543/jcasbin-sample"
    val dbUsername = "postgres"
    val dbPassword = "postgres"

    val modelText = object {}.javaClass.getResource("../model.conf")?.readText()
    val model = newModel(modelText)
    val adapter = JDBCAdapter(dbDriver, dbUrl, dbUsername, dbPassword);

    return Enforcer(model, adapter, true)
}