const mongoDb = require('mongodb')
const dbName = 'sampleauth'
const DB_USER = process.env.DB_USER
const DB_PASS = process.env.DB_PASS
const url = `mongodb+srv://${DB_USER}:${DB_PASS}@cluster0.lowctzs.mongodb.net/?retryWrites=true&w=majority`
const mongoClient = mongoDb.MongoClient



module.exports = {mongoDb, dbName, url, mongoClient}