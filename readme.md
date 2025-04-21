```
 __  __               _                          __  __         _  _             
|  \/  |  ___  _ __  | |_   ___   ___           |  \/  |  __ _ (_)| |  ___  _ __ 
| |\/| | / _ \| '_ \ | __| / _ \ / _ \  _____   | |\/| | / _` || || | / _ \| '__|
| |  | ||  __/| | | || |_ |  __/|  __/ |_____|  | |  | || (_| || || ||  __/| |   
|_|  |_| \___||_| |_| \__| \___| \___|          |_|  |_| \__,_||_||_| \___||_|   
                                                                                 

                         
```


# Mentee Selection Mailer
A Go tool to automate sending personalized HTML and plaintext emails to your mentees. Just update the config and CSV, and you're good to go.


Used by `Brain & Cognitive Science Club` and 
`Electronics Club` at IIT Kanpur to onboard their summer mentees.

---

## Setup

### 1. Update the `config.yml` file

Make sure to configure the following values in the `config.yml` file:

- **Mail Configuration**: Update the SMTP server details (host, port, user, and password) under `MAIL`.
- **Club Information**: Enter your club's details like name, site, email, and coordinators' names and emails.
- **Project Information**: Update the project name, WhatsApp link, and the mentors' names.
- **Template Text**: Customize the email template text to your liking.

### 2. Prepare the `sample.csv` File

This file should contain the names and Roll numbers of all mentees. Example format:

```csv
Name,RollNo
Akshat,230094
...
```

Do not add "@iitk.ac.in" to the Roll number or username, as it's already added in the code.

### 3. Run the App

Once the `config.yml` and `sample.csv` files are updated, you are ready to send the emails. 

In your terminal, run the following command:

```bash
go run .
```

**Note**: It is advisable to first send test emails to yourself to ensure everything works as expected.

This will send personalized emails (both HTML and plaintext) to all mentees listed in the `sample.csv` file.

---

## Enjoy!

That's it! You’re all set to email your mentees. Let me know if you need any tweaks or have further questions.
