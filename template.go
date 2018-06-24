package main

const entryTmpl = `
entity "{{ .Name }}" {
{{- if .Comment.Valid }}
  {{ .Comment.String }}
  ..
{{- end }}
{{- range .Columns }}
  {{- if .IsPrimaryKey }}
  {{ printf "+ [PK] %-33s" .Name }} {{ printf "\t\t%s" .DataType }} {{- if .Comment.Valid }} : {{ .Comment.String }}{{- end }}
  {{- end }}
{{- end }}
  --
{{- range .Columns }}
  {{- if not .IsPrimaryKey }}
  {{ printf "%-40s" .Name }} {{ printf "\t\t%s" .DataType }} {{- if .Comment.Valid }} : {{ .Comment.String }}{{- end }}
  {{- end }}
{{- end }}
}
`

const relationTmpl = `
{{ if .IsOneToOne }} {{ .SourceTableName }} ||-|| {{ .TargetTableName }}{{else}} {{ .SourceTableName }} }-- {{ .TargetTableName }}{{end}}
`
