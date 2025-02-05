# awsips - AWS IP Ranges CLI ☁️

awsips is a simple, fast, and efficient CLI tool for querying AWS IP ranges by region and service.

### **📦 Build Instructions**


**Prerequisites**

Go 1.20+ installed (Download Go)

Git installed

**Clone and Build**

`git clone https://github.com/jordlevy/awsips.git`

`cd awsips`

`go build -o awsips`

This creates an awsips binary in the current directory.

**Install Globally**

To install awsips system-wide:

`sudo mv awsips /usr/local/bin/`

Now you can run awsips from anywhere in your terminal.

### **🛠 Usage**

awsips [OPTIONS]

### **📌 Basic Commands**

| Command                                                                     | Description                                          |
| --------------------------------------------------------------------------- | ---------------------------------------------------- |
| `awsips --list` or `awsips -l`                                              | List all available AWS regions and services.         |
| `awsips --region REGION` or `awsips -r REGION`                              | Show all IP ranges for a specific region (sorted).   |
| `awsips --service SERVICE` or `awsips -s SERVICE`                           | Show all IP ranges for a specific service (sorted).  |
| `awsips --region REGION --service SERVICE` or `awsips -r REGION -s SERVICE` | Show IPs for a service in a region.                  |
| `awsips -l -r REGION`                                                       | List all services available in a given region.       |
| `awsips -l -s SERVICE`                                                      | List all regions where a given service is available. |
| `awsips --help` or `awsips -h`                                              | Show this help menu.                                 |

**⚙️ Configuration (.awsipsrc)**

You can customize the AWS IP ranges source by creating a .awsipsrc file in your home directory.

Example (~/.awsipsrc):

{
"aws_ip_ranges_url": "https://my-custom-source.com/ip-ranges.json"
}

If this file is missing, awsips defaults to:

https://ip-ranges.amazonaws.com/ip-ranges.json

Some might specifically want to use a custom source JSON, so that's why I opted to make this configurable.

### **📜 License**

This project is licensed under the MIT License. See LICENSE for details.

### **⚠️ Disclaimer**

This tool is not affiliated with or endorsed by AWS.

The IP ranges are sourced from AWS's official JSON file.

Always verify IPs before making firewall changes.

### **🚀 Contributions**

PRs and feature requests are welcome! Feel free to submit an issue or fork the project. I might need to do some work on Windows.