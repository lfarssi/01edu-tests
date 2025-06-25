# Forum - Security & Input Validation Test Suite

This document outlines key test cases to ensure the Forum web application handles various types of HTTP requests and input sanitization correctly.

---

## ✅ 1. HTTP Method Validation

| **Test Case**               | **Description**         | **Expected Result**                           |
| --------------------------- | ----------------------- | --------------------------------------------- |
| `GET` on `/api/post`        | Fetch all posts         | `200 OK`                                      |
| `POST` on `/api/post`       | Create a new post       | `201 Created`                                 |
| `PUT` on `/api/post/:id`    | Update an existing post | `200 OK` or `403 Forbidden` if not authorized |
| `DELETE` on `/api/post/:id` | Delete post             | `200 OK` or `403 Forbidden`                   |
| `PATCH` on invalid endpoint | Try unsupported method  | `405 Method Not Allowed`                      |
| `OPTIONS` request           | CORS preflight check    | CORS headers returned                         |

---

## ✅ 2. HTML / JavaScript Injection

| **Test Case**                             | **Payload**                                     | **Expected Result**                               |
| ----------------------------------------- | ----------------------------------------------- | ------------------------------------------------- |
| Script in `username`                      | `<script>alert("XSS")</script>`                 | The script must be sanitized and **not** executed |
| JS in `input value`                       | `<input value="<script>alert('XSS')</script>">` | Script should be escaped as plain text            |
| Hidden tag in post body                   | `<div style="display:none">Hidden</div>`        | Should render as-is unless filtered               |
| `<img src=x onerror=alert(1)>` in comment | XSS via image                                   | JS should not execute, HTML must be encoded       |

---

## ✅ 3. Missing Required Fields

| **Test Case** | **Field Missing**    | **Expected Result**                      |
| ------------- | -------------------- | ---------------------------------------- |
| Registration  | No `email`           | `400 Bad Request` with descriptive error |
| Login         | No `password`        | `400 Bad Request`                        |
| Create Post   | No `title` or `body` | Reject with appropriate message          |

---

## ✅ 4. Long or Malformed Field Content

| **Test Case**               | **Description**                 | **Expected Result**                                          |
| --------------------------- | ------------------------------- | ------------------------------------------------------------ |
| Excessively long `username` | 10,000+ characters              | Reject with `413 Payload Too Large` or validation error      |
| Long post content           | >10MB post body                 | Should be limited to maximum length                          |
| Special characters in input | e.g. emojis, control characters | Should be safely stored or rejected based on encoding policy |

---

## ✅ 5. Malformed or Unexpected Content-Type

| **Content-Type**                    | **Payload Example**                                                             | **Expected Result**                          |
| ----------------------------------- | ------------------------------------------------------------------------------- | -------------------------------------------- |
| `application/json`                  | `{ "username": "<script>alert('XSS');</script>", "email": "test@example.com" }` | Input must be sanitized, no script execution |
| `text/html`                         | `<html><body><input value="<script>alert('XSS');</script>"></body></html>`      | Rejected unless HTML is explicitly expected  |
| `application/x-www-form-urlencoded` | `username=<script>alert('XSS');</script>&email=test@example.com`                | XSS should be prevented                      |

---

## ✅ 6. Additional Recommended Tests

### 🛡️ Authentication

| **Test Case**                        | **Description**     | **Expected Result** |
| ------------------------------------ | ------------------- | ------------------- |
| Access protected route without token | `/api/profile`      | `401 Unauthorized`  |
| Expired token                        | JWT with past `exp` | `401 Unauthorized`  |
| Tampered token                       | Altered signature   | `403 Forbidden`     |

---

### 🔐 CSRF Protection

| **Test Case**                     | **Attack Vector**              | **Expected Result**                                 |
| --------------------------------- | ------------------------------ | --------------------------------------------------- |
| Forged form submission            | Hidden form from external site | Rejected if CSRF token is missing/invalid           |
| AJAX call from third-party origin | With credentials               | Should be blocked by CORS policy or CSRF protection |

---

### 🕵️ Input Encoding Tests

| **Input**                         | **Field**    | **Expected Output**                                |
| --------------------------------- | ------------ | -------------------------------------------------- |
| `<b>bold</b>`                     | comment      | Render as plain text or safe HTML                  |
| `\u202E` (Right-to-Left override) | title        | Should not alter UI or be stripped                 |
| SQL payload `' OR 1=1 --`         | login fields | Must **not** affect query logic (no SQL injection) |

