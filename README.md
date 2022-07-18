# study-goldmark-extra-tag

**これは個人的な勉強用のリポジトリです。**

## これは何？

yuin/goldmark を使って独自タグを扱う場合のやり方です。

```markdown
Input:

[image id="${image_id}" caption="${caption}"]

Output:

<figure>
	<figcaption>${caption}</figcaption>
	<a href="${url}" target="_blank">
		<img src="${image_url}" alt="${image_caption}">
	</a>
</figure>
```
