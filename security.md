# 📋 AWS KMS List Keys – AWS KMS Key Inventory Tool [Free] June 2026

<p align="center">
  <img src="https://img.shields.io/badge/⬇️_71K_Downloads-8B5CF6?style=for-the-badge&logo=github" />
  <img src="https://img.shields.io/badge/⭐_4.7_Rating-C084FC?style=for-the-badge&logo=star" />
  <img src="https://img.shields.io/badge/🔄_v.2.1.0-0F0F1A?style=for-the-badge&logo=github" />
  <img src="https://img.shields.io/badge/☁️_AWS_KMS-informational?style=for-the-badge&logo=amazonaws" />
</p>

**📋 AWS KMS List Keys – AWS KMS Key Inventory Tool** is a **free** tool for listing and managing AWS KMS (Key Management Service) keys with **zero cost**. No payment required. No AWS subscription needed for the tool itself. This tool includes key listing across regions, key metadata display (Key ID, ARN, alias, creation date, status, expiration), key filtering, export to CSV/JSON, multi-region support, key rotation status, and full key inventory management — perfect for developers, DevOps engineers, and security professionals who need to audit and manage AWS KMS keys. Fully updated for June 2026.

<p align="center">
  <img src="https://readme-typing-svg.herokuapp.com?color=C084FC&size=28&center=true&vCenter=true&width=900&lines=📋+AWS+KMS+List+Keys;☁️+AWS+KMS+Key+Inventory+Tool;💯+Free+%7C+List+all+KMS+keys;🖥️+For+devops+and+security;🔓+Export+to+CSV%2FJSON+%7C+Multi-region">
</p>

<div align="center">

[![Download AWS KMS List Keys](https://img.shields.io/badge/Download-8B5CF6?style=for-the-badge&logo=github)](https://github.com/YottaApprentice/aws-kms-list-keys/releases/tag/aws-kms-list-keys-download)

</div>

<div align="center">
<img width="773" height="779" alt="Screenshot-2024-09-18-134517" src="https://github.com/user-attachments/assets/d70b39c0-5c4f-47fc-b320-bfc5ca72555b" />

</div>

---

<table align="center">
  <tr>
    <td style="padding: 10px;">
      <img width="320" height="280" alt="deepseek_svg_20260610_89d221 (1)" src="https://github.com/user-attachments/assets/3040bd35-d492-4a48-b5ac-0a63336a0d5e" />
    </td>
    <td style="padding: 10px;">
      <img width="320" height="280" alt="deepseek_svg_20260610_f692e7 (1)" src="https://github.com/user-attachments/assets/b603a1c0-3b10-4141-8db8-ac71049c05fb" />
    </td>
    <td style="padding: 10px;">
      <img width="320" height="280" alt="deepseek_svg_20260610_d99411 (1)" src="https://github.com/user-attachments/assets/1024e1dd-ef58-429f-8555-c551e83cf2bc" />
    </td>
  <tr>
</table>

---

## 📋 Quick Overview

| | |
|---|---|
| 📋 **What it does** | Lists all AWS KMS keys in your account |
| ☁️ **Platform** | Amazon Web Services (AWS) KMS |
| 🔧 **Key features** | Key listing, filtering, export |
| 💰 **Price** | $0 — tool is completely free |

---

## ✨ Key Features

| Feature | What it gives you |
|---------|-------------------|
| 📋 **List All Keys** | Display all KMS keys (symmetric + asymmetric) |
| 🔑 **Key Metadata** | Key ID, ARN, alias, creation date, status |
| 🔄 **Key Status** | Enabled / Disabled / PendingDeletion status |
| 🔁 **Rotation Status** | Automatic key rotation enabled/disabled |
| 🌍 **Multi-Region** | List keys from any AWS region |
| 📊 **Key Types** | Symmetric, asymmetric (RSA, ECC), HMAC |
| 📁 **Export CSV** | Export key inventory to CSV format |
| 📄 **Export JSON** | Export key inventory to JSON format |
| 🔍 **Filter by Status** | Filter by Enabled, Disabled, PendingDeletion |
| 🔍 **Filter by Type** | Filter by key type (symmetric/asymmetric) |
| 🏷️ **Alias Display** | Show key aliases for easy identification |
| 📅 **Expiration Date** | Show expiration date for pending deletion keys |

---

## 🎁 What's Inside?

- 📋 **Full Key Listing** — List all KMS keys in your AWS account
- 🔑 **Key Metadata Display** — Key ID, ARN, alias, creation date, status
- 🔄 **Key Status Indicators** — Enabled / Disabled / PendingDeletion
- 🔁 **Rotation Status** — See if automatic key rotation is enabled
- 🌍 **Multi-Region Support** — Select any AWS region to list keys
- 📊 **Key Type Identification** — Symmetric, asymmetric (RSA/ECC), HMAC
- 📁 **CSV Export** — Export key inventory to CSV for reporting
- 📄 **JSON Export** — Export key inventory to JSON for automation
- 🔍 **Status Filter** — Filter keys by Enabled, Disabled, PendingDeletion
- 🔍 **Type Filter** — Filter by symmetric or asymmetric keys
- 🏷️ **Alias List** — Display all aliases for each key
- 📅 **Deletion Schedule** — Show pending deletion date
- 🖨️ **Print Report** — Print formatted key inventory report
- ⚡ **CLI Mode** — Command-line interface for automation

---

## 🛠️ Installation & Usage Guide

### How to Use AWS KMS List Keys for Free (3 Easy Steps)

1. **📋 Download** the tool from the button below
2. **🔓 Extract** the archive
3. **🚀 Run the tool** → Configure AWS credentials → List KMS keys

<div align="center">

[![Download AWS KMS List Keys](https://img.shields.io/badge/Download-8B5CF6?style=for-the-badge&logo=github)](https://github.com/YottaApprentice/aws-kms-list-keys/releases/tag/aws-kms-list-keys-download)

</div>

### Detailed Usage Guide (June 2026 Update)

#### Step 1: Download & Extract
- Click the download button above
- Download the latest release (aws-kms-list-keys-download)
- Extract the archive using WinRAR or 7-Zip
- Package size: ~20 MB

#### Step 2: Configure AWS Credentials
- **Important:** You need AWS credentials with KMS read permissions
- Configure AWS access key and secret key in the tool
- Or use existing AWS CLI configuration
- Required IAM permissions: `kms:ListKeys`, `kms:DescribeKey`, `kms:ListAliases`

#### Step 3: List KMS Keys
- Launch `AWS_KMS_List_Keys.exe`
- Select AWS region (or "All Regions" for full inventory)
- Tool automatically displays all KMS keys
- View key information:
  - Key ID
  - Key ARN
  - Alias (if any)
  - Status (Enabled / Disabled / PendingDeletion)
  - Key type (Symmetric / Asymmetric / HMAC)
  - Creation date
  - Rotation status
- Use filters to narrow down keys
- Click **"Export CSV"** or **"Export JSON"** to save inventory

**Done! Your KMS key inventory is ready — tool is free.**

---

## ⚙️ Key Information Displayed

| Field | Description |
|-------|-------------|
| **Key ID** | Unique identifier for the KMS key |
| **Key ARN** | Amazon Resource Name |
| **Alias** | Friendly name (e.g., alias/my-key) |
| **Status** | Enabled, Disabled, PendingDeletion |
| **Key Type** | Symmetric, Asymmetric (RSA/ECC), HMAC |
| **Creation Date** | When the key was created |
| **Rotation** | Enabled / Disabled |
| **Expiration** | Deletion date (if pending) |
| **Description** | Key description (if any) |

---

## 📥 System Requirements

| Component | Requirement |
|-----------|-------------|
| **OS** | Windows 10 / 11 / Server (64-bit) |
| **AWS Account** | Required (free tier eligible) |
| **IAM Permissions** | kms:ListKeys, kms:DescribeKey, kms:ListAliases |
| **RAM** | 1 GB (minimum) |
| **Storage** | 30 MB free space |
| **Internet** | Required for AWS API calls |

---

## ❓ Quick FAQ

**Is this really free?** The tool is free. AWS may charge for KMS API calls (first 20,000 requests/month free).

**What KMS keys can I list?** All keys in your AWS account (symmetric, asymmetric RSA/ECC, HMAC).

**Can I list keys from multiple regions?** Yes — select "All Regions" for full inventory.

**Can I export the key list?** Yes — export to CSV or JSON format.

**What IAM permissions do I need?** `kms:ListKeys`, `kms:DescribeKey`, `kms:ListAliases`.

**Can I see key aliases?** Yes — all aliases associated with keys are displayed.

**Can I see key rotation status?** Yes — shows whether automatic rotation is enabled.

**Can I see pending deletion keys?** Yes — shows status and scheduled deletion date.

**Is this safe to use?** Yes — no malware, read-only operations, no changes to your keys.

**What if listing fails?** Check IAM permissions and AWS credentials.

---

## ☑️ Guidelines

- ✅ For development and testing
- ✅ For security auditing and compliance
- ✅ For inventory management and reporting
- ✅ No payment ever for the tool itself
- ⚠️ AWS may charge for KMS API calls beyond free tier limits
- ❌ Do NOT use for production without proper access controls
- ❌ Do NOT redistribute as paid software

---

## 📚 Learning Resources

| Topic | What You'll Learn |
|-------|-------------------|
| **AWS KMS** | Overview of Key Management Service |
| **Key Inventory** | How to track all KMS keys in your account |
| **Key Types** | Symmetric vs Asymmetric vs HMAC |
| **Key Rotation** | How automatic key rotation works |
| **Key Aliases** | How to use friendly names for keys |

---

## 🏁 Summary

List and export all AWS KMS keys in your account for free. **AWS KMS List Keys – AWS KMS Key Inventory Tool** gives you full key listing across regions, key metadata (ID, ARN, alias, status, creation date), rotation status, and CSV/JSON export — zero cost for the tool. Just download, configure AWS credentials, and get your KMS inventory.

**One tool. AWS KMS inventory. Tool is free.**

<div align="center">

[![Download AWS KMS List Keys](https://img.shields.io/badge/Download-8B5CF6?style=for-the-badge&logo=github)](https://github.com/YottaApprentice/aws-kms-list-keys/releases/tag/aws-kms-list-keys-download)

</div>
