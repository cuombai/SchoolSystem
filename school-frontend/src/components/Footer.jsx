import React from "react";

const Footer = () => {
  return (
    <footer style={styles.footer}>
      <div style={styles.section}>
        <h4>Contact Us</h4>
        <p>Email: info@kibali-school.ac.ke</p>
        <p>Phone: +254 712 345678</p>
        <p>Location: Kisumu, Kenya</p>
      </div>

      <div style={styles.section}>
        <h4>Quick Links</h4>
        <ul style={styles.links}>
          <li><a href="/">Home</a></li>
          <li><a href="/about">About Us</a></li>
          <li><a href="/contact">Contact</a></li>
          <li><a href="/login">Login</a></li>
        </ul>
      </div>

      <div style={styles.section}>
        <h4>Follow Us</h4>
        <div style={styles.socials}>
          <a href="https://facebook.com" target="_blank" rel="noreferrer">Facebook</a>
          <a href="https://twitter.com" target="_blank" rel="noreferrer">Twitter</a>
          <a href="https://instagram.com" target="_blank" rel="noreferrer">Instagram</a>
        </div>
      </div>
    </footer>
  );
};

const styles = {
  footer: {
    display: "flex",
    justifyContent: "space-around",
    backgroundColor: "#00264d",
    color: "#fff",
    padding: "2rem",
    marginTop: "2rem",
  },
  section: {
    flex: 1,
    padding: "0 1rem",
  },
  links: {
    listStyle: "none",
    padding: 0,
  },
  socials: {
    display: "flex",
    flexDirection: "column",
    gap: "0.5rem",
  },
};

export default Footer;
