import React from "react";
//import './ProfilePage.css';
import styles from "./profilePage.module.css";

const ProfilePage = () => {
  return (
    <div>
      <header className={styles.profileHeader}>
        <button className={styles.backButton}>←</button>
        <h5>Profile</h5>
        <img
          src="profile-pic-url"
          //alt="Profile"
          className={styles.profilePic}
        />
        <h2>Jenny, 22</h2>
      </header>
      {/* <div className={styles.profileInfo}>
        <div className={styles.profileNameEdit}>
          <h2>Jenny, 22</h2>
          <button className={styles.editButton}>✎</button>
        </div>
      </div> */}

      <div className={styles.settings}>
        <div className={styles.section}>
          <h3>Account Settings</h3>
          <button className={styles.editLink}>Edit</button>
        </div>
        <div className={styles.field}>
          <label>Name</label>
          <p>Jenny</p>
        </div>
        <div className={styles.field}>
          <label>Phone Number</label>
          <p>+91 9876543210</p>
        </div>
        <div className={styles.field}>
          <label>Date of Birth</label>
          <p>02-05-1997</p>
        </div>
        <div className={styles.field}>
          <label>Email</label>
          <p>abcqwertyu@gmail.com</p>
        </div>

        <div className="section">
          <h3>Plan Settings</h3>
        </div>
        <div className={styles.field}>
          <label>Current Plan</label>
          <p>Free</p>
        </div>

        <div className="section">
          <h3>Discovery Settings</h3>
        </div>
        <div className={styles.field}>
          <label>Location</label>
          <p>My Current Location</p>
        </div>
        <div className={styles.field}>
          <label>Preferred Languages</label>
          <p>English</p>
        </div>
        <div className={styles.field}>
          <label>Show Me</label>
          <p>Men</p>
        </div>
        <div className={styles.sliderField}>
          <label>Age Range</label>
          <input type="range" min="18" max="60" value="22-34" />
          <p>22 - 34</p>
        </div>
        <div className={styles.sliderField}>
          <label>Maximum Distance</label>
          <input type="range" min="0" max="100" value="100" />
          <p>100km</p>
        </div>
      </div>

      <button className={styles.logoutButton}>Logout</button>
      <button className={styles.deleteAccountButton}>Delete Account</button>

      <footer className={styles.navbar}>
        <button>Home</button>
        <button>Likes</button>
        <button>Messages</button>
        <button>Profile</button>
      </footer>
    </div>
  );
};

export default ProfilePage;
