import "dotenv/config"
import readline from "node:readline"
import { faker } from "@faker-js/faker"
import * as nodemailer from "nodemailer"

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
})

const transporter = nodemailer.createTransport({
    host: "localhost",
    port: 1025,
    secure: false,
    auth: undefined,
    newline: "unix",
})

function questionPromise(query: string): Promise<string> {
    return new Promise((resolve) => {
        rl.question(query + ": ", (answer) => {
            resolve(answer)
        })
    })
}

async function generateFake() {
    let to = await questionPromise("Email")
    const amountString = await questionPromise("Amount")
    const amount = parseInt(amountString, 10)

    if (!to.includes("@")) {
        to += "@localhost"
    }

    for (let i = 0; i < amount; i++) {
        const firstName = faker.person.firstName()
        const lastName = faker.person.lastName()
        const fakeEmail = faker.internet.email({ firstName, lastName })

        const subject = faker.lorem.sentence()
        const text = faker.lorem.paragraph()
        const sent = faker.date.past().toISOString()

        const info = await transporter.sendMail({
            from: `${firstName} ${lastName} <${fakeEmail}>`, // Visible From: header
            to, // Visible To: header
            sent, // Visible Date: header
            envelope: {
                from: fakeEmail, // Actual MAIL FROM (for bounces)
                to: [to, fakeEmail], // Actual RCPT TO (for bounces)
            },
            subject,
            text,
        })

        console.log(info)
    }

    rl.close()
}

generateFake()
