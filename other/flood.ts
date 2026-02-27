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

function questionPromise(query: string, defaultAnswer: string): Promise<string> {
    const noAsk = process.argv.includes("--no-ask")

    return new Promise((resolve) => {
        if (noAsk) {
            resolve(defaultAnswer)
        } else {
            rl.question(query + ": ", (answer) => {
                resolve(answer || defaultAnswer)
            })
        }
    })
}

async function generateFake() {
    let to = await questionPromise("Email", "test@localhost")
    let toName = to.split("@")[0]
    const amountString = await questionPromise("Amount", "50")
    const amount = parseInt(amountString, 10)

    if (!to.includes("@")) {
        toName = to
        to += "@localhost"
    }

    for (let i = 0; i < amount; i++) {
        const firstName = faker.person.firstName()
        const lastName = faker.person.lastName()
        const fakeEmail = faker.internet.email({ firstName, lastName })

        const subject = faker.lorem.sentence()
        const sent = faker.date.past().toISOString()
        
        const randomBody = faker.lorem.paragraphs({ min: 1, max: 20 })
            .split('\n')
            .map((p) => `<p>${faker.lorem.words({ min: p.split(' ').length, max: p.split(' ').length * 6 })}</p>`)
            .join('')
        const toNameCapitalized = toName.charAt(0).toUpperCase() + toName.slice(1)
        const text = `<p>Dear ${toNameCapitalized},</p>${randomBody} <p>Best Regards,<br>${firstName} ${lastName}</p>`

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
