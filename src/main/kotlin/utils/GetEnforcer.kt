package utils

import com.mongodb.client.MongoClients
import org.casbin.adapter.JDBCAdapter
import org.casbin.jcasbin.main.CoreEnforcer.newModel
import org.casbin.jcasbin.main.Enforcer
import org.casbin.jcasbin.model.Model
import org.casbin.jcasbin.persist.Adapter
import org.jim.jcasbin.MongoAdapter

fun getEnforcer(casbinModel: Model?, filteredAdapter: Adapter?): Enforcer {
    val adapter = filteredAdapter ?: getAdapter()
    val model = casbinModel ?: getModel()

    return Enforcer(model, adapter, true)
}

fun getModel(): Model? {
    val modelText = object {}.javaClass.getResource("/model.conf")?.readText()
    val model = newModel(modelText)
    return model
}

private fun getAdapter(): Adapter {
    val adapterType = System.getenv("DB_TYPE")?.lowercase() ?: "postgres"
    if (adapterType == "postgres")
        return getPostgresAdapter()
    return getMongoAdapter()
}

fun getPostgresAdapter(): Adapter {
    val dbHost = System.getenv("DB_HOST") ?: "localhost"
    val dbPort = System.getenv("DB_PORT") ?: "6543"
    val dbDriver = "org.postgresql.Driver"
    val dbUrl = "jdbc:postgresql://$dbHost:$dbPort/jcasbin-sample"
    val dbUsername = "db-user"
    val dbPassword = "db-password"

    return JDBCAdapter(dbDriver, dbUrl, dbUsername, dbPassword)
}

private fun getMongoAdapter(): Adapter {
    val dbHost = System.getenv("MONGO_DB_HOST") ?: "localhost"
    val dbPort = System.getenv("MONGO_DB_PORT") ?: "27017"
    val dbUsername = "db-user"
    val dbPassword = "db-password"
    val mongoClient = MongoClients.create("mongodb://$dbUsername:$dbPassword@$dbHost:$dbPort")
    val adapter = MongoAdapter(mongoClient, "jcasbin-sample");
    return adapter
}