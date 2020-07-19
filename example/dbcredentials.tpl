
production:
  adapter: postgresql
  database: {{ .Database }}
  host: {{.DatabaseHost }}
  username: {{ .DatabaseUsername }}
  password: {{ .DatabasePassword }}
  pool: 5
