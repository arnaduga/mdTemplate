# ORR for {{ .applicationName }}

Author : [{{ .author.name }}](mailto:{{ .author.email }})



{{ range .sections }}
## Section: {{ .title }}

{{ range .criterias }}
 
![img]({{ .status }}.svg): {{ .name }}
{{ range .attachments }}
- [{{ .name }}]({{ .file }})
{{ end }}
  
---

{{ end }}
{{ end }}