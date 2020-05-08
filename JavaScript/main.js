var admin = require("firebase-admin");

// Fetch the service account key JSON file contents
var serviceAccount = require("/Users/stefan/.secret/hybrid-cloud-22365-firebase-adminsdk-ca37q-d1e808e47b.json");

// Initialize the app with a service account, granting admin privileges
admin.initializeApp({
    credential: admin.credential.cert(serviceAccount),
    databaseURL: "https://hybrid-cloud-22365.firebaseio.com"
});

// As an admin, the app has access to read and write all data, regardless of Security Rules
var db = admin.database();
var ref = db.ref("/");
ref.once("value", function(snapshot) {
    console.log(snapshot.val());
});