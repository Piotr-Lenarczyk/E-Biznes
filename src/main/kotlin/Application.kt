package com.example

import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.http.*
import io.ktor.serialization.kotlinx.json.*
import io.ktor.server.application.*
import io.ktor.server.plugins.contentnegotiation.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.serialization.json.Json

fun Application.module() {
    // Install the ContentNegotiation feature for the server-side
    install(ContentNegotiation) {
        json(Json {
            prettyPrint = true
            isLenient = true
        })
    }

    val client = HttpClient(CIO) {
        // Configure the client to send requests
    }

    routing {
        route("/send-discord-message") {
            post {
                val message = call.receive<String>() // Get message from request body
                sendDiscordMessage(client, message) // Send message to Discord
                call.respondText("Message sent to Discord!") // Corrected: No 'typeInfo' parameter
            }
        }
    }
}

suspend fun sendDiscordMessage(client: HttpClient, message: String) {
    val webhookUrl = "${WEBHOOK_URL}"

    // Remove any surrounding quotes from the message (in case it has extra quotes)
    val cleanMessage = message.trim('"')

    val payload = """
        {
            "content": "$cleanMessage"
        }
    """.trimIndent()
    println(payload)

    try {
        val response = client.post(webhookUrl) {
            contentType(ContentType.Application.Json)
            setBody(payload)
        }

        println("Discord Response: ${response.status}")
    } catch (e: Exception) {
        println("Error sending to Discord: ${e.message}")
    }
}

