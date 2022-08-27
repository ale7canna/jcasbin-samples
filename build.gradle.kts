import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    kotlin("jvm") version "1.7.0"
    kotlin("plugin.allopen") version "1.7.0"
    id("org.jetbrains.kotlinx.benchmark")  version "0.4.4"
    application
}

configure<org.jetbrains.kotlin.allopen.gradle.AllOpenExtension> {
    annotation("org.openjdk.jmh.annotations.State")
}

group = "me.user"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
    maven {
        url = uri("https://jitpack.io")
    }
}

dependencies {
    implementation("org.casbin:jcasbin:1.24.2")
    implementation("org.casbin:jdbc-adapter:2.3.1")
    implementation("com.github.zzl221000:jcasbin-mongo-adapter:v1.0")
    implementation("org.mongodb:mongodb-driver-sync:4.6.1")
    implementation("org.jetbrains.kotlinx:kotlinx-benchmark-runtime:0.4.4")
    testImplementation(kotlin("test"))
}

tasks.test {
    useJUnitPlatform()
}

tasks.withType<KotlinCompile> {
    kotlinOptions.jvmTarget = "1.8"
}

application {
    mainClass.set("MainKt")
}

tasks.jar {
    manifest {
        attributes["Main-Class"] = "MainKt"
    }
    duplicatesStrategy = DuplicatesStrategy.EXCLUDE
    configurations.compileClasspath.get()
        .forEach {
            from(if (it.isDirectory) it else zipTree(it))
                .exclude("META-INF/**")
                .exclude("META-INF/*.DSA")
                .exclude("META-INF/*.RSA")
        }
}

benchmark {
    targets {
        register("main")
    }
}