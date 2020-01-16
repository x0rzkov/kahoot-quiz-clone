const Instagram = require('instagram-web-api')

const { username, password } = process.env

const client = new Instagram({ username, password })

;(async () => {
  await client.login()
  const profile = await client.getProfile()

  console.log(profile)
})()
