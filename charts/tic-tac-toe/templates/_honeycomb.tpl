{{ define "templates.honeycomb.env" -}}
{{ if hasKey .Values "honeycomb" -}}
- name: HONEYCOMB_API_KEY
  valueFrom:
    secretKeyRef:
      name: honeycomb-secret
      key: api_key
- name: HONEYCOMB_DATASET
  value: {{ .Values.honeycomb.dataset }}
{{- end}}
{{- end }}