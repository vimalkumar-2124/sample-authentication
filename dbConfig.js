const mongoDb = require('mongodb')
const dbName = 'users'
const url = `mongodb://localhost:27017/${dbName}`
const mongoClient = mongoDb.MongoClient



module.exports = {mongoDb, dbName, url, mongoClient}