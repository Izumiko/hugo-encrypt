---
title: "Example"
date: 2019-07-22
tags: ["go", "golang", "hugo", "example"]
draft: false
---

#Content before encryption.

<!--more-->
{{% hugo-encrypt "PASSWORD" %}}
##Encrypted content.
{{% /hugo-encrypt %}}