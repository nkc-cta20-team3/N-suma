// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyAjnRKiJvyN9bdb1FFjPlRcAyaAdgAaZEc",
  authDomain: "n-suma.firebaseapp.com",
  projectId: "n-suma",
  storageBucket: "n-suma.appspot.com",
  messagingSenderId: "808452019699",
  appId: "1:808452019699:web:deb3619ab2626ba1e96e86",
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

export { app };
