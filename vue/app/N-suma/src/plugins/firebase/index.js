/**
 * firebase/index.js
 * included in `@store/index.js`
 */

// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth, GoogleAuthProvider } from 'firebase/auth'
import { getFirestore } from 'firebase/firestore'
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// Initialize Firebase
const app = initializeApp({
  apiKey: "AIzaSyAjnRKiJvyN9bdb1FFjPlRcAyaAdgAaZEc",
  authDomain: "n-suma.firebaseapp.com",
  projectId: "n-suma",
  storageBucket: "n-suma.appspot.com",
  messagingSenderId: "808452019699",
  appId: "1:808452019699:web:deb3619ab2626ba1e96e86",
});

export const auth = getAuth(app)
export const db = getFirestore(app)
export const provider = new GoogleAuthProvider()