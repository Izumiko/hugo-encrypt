# Hugo Encrypt

**Hugo-Encrypt** is a golang port of [Hugo Encryptor](https://github.com/Li4n0/hugo_encryptor)
   
**Hugo-Encrypt** is a tool to protect your [Hugo](https://gohugo.io) posts. It uses AES-256-GCM to encrypt the contents of your posts, and inserts a snippet of `<script>` code to verify whether the password is correct or not in readers' browser. Without a correct key, nobody can decrypt your private posts.

## Installation

Environmental dependence: go 1.11+

**Step 1: Install Hugo-Encrypt.**

    $ go get github.com/Izumiko/hugo_encrypt.git

**Step 2: Place `hugo-encrypt` in the root directory of your blog, or add `hugo-encrypt` to `PATH`**

    $ cp hugo-encrypt /path/to/your/blog/

    or

    $ export PATH=/path/of/hugo-encrypt/:$PATH

**Step 3: Place `shortcodes/hugo-encryptor.html` in the shortcode directory of your blog:**

    $ mkdir -p /path/to/your/blog/layouts/shortcodes
    $ cp shortcodes/hugo-encryptor.html /path/to/your/blog/layouts/shortcodes/hugo-encryptor.html

**Step 4: Merge i18n translation files and/or add your own language.**

    $ cp -r i18n /path/to/your/blog/

## Usage

**Step 1: Use `hugo-encryptor` tag surround the text you want to encrypt **

**Attention! There must be some text and the `<!--more-->` tag before the hugo-encryptor:**

```markdown
---
title: "This Is An Encrypted Post"
---

**There must be some text, and the summary tag is also needed:**
<!--more-->
{{% hugo-encryptor "PASSWORD" %}}

# You cannot see me unless you've got the password!

This is the content you want to encrypt!

**Do remember to close the `hugo-encryptor` shortcodes tag:**

{{% /hugo-encryptor %}}
```

**Step 2: Generate your site as usual**

It may be something like:

    $ hugo

**Step 3: Get the encryption done!**

    $ ./hugo-encrypt

    or (if added to PATH)

    $ hugo-encrypt


Then all the private posts in your `public` directory would be encrypted thoroughly, congrats!

## Configuration

Although the **Hugo-Encrypt** can run without any configure, we provide some settings params to help you configure **Hugo-Encrypt** to your liking.

### Language

**Hugo-Encrypt** uses i18n settings to display. You can change it by param below. Be sure to add the corresponding language file to the i18n folder.

```toml
[params]
 		 ......
  DefaultContentLanguage = "zh-cn"
```

### The way of client password storage

As default,**Hugo-Encrypt** use `localStorage` to storage the password in client. By adding `hugoEncryptorStorage` param in your blog's config file, you can change the storage method into `sessionStorage`. Such as below:

```toml
[params]
 		 ......
  hugoEncryptorStorage = "session" # or "local"
```

For the difference of two storage ways:

* **localStorage**:

  Once a reader input the correct password,the authorization status will not expire, the reader can read the article at any time without having to enter the password again. Unless you change the password or the reader clean his browser cache.

* **sessionStorage**:

  If a reader input the correct password, he could read the article without having to enter the password again until the user close his browser.

### Style

As default, **Hugo-Encrypt** has no style,but we have already give all the visual element a class name, you can add style for them in your css files.

## Attentions

* Do remember to keep the source code of your encrypted posts private. Never push your blog directory into a public repository.

* Every time when you generate your site, you should run `$ hugo-encrypt` again to encrypt the posts which you want to be protected. If you are worried about you will forgot that, it's a good idea to use a shell script to take the place of  `$ hugo` ,such as below:

  ```bash
  #!/bin/bash
  hugo && hugo-encrypt
  ```

  

