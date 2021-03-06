# Nmap Issue
Issue generated by nmap module.

## Addresses
| Address | Type | Vendor |
| ------- | ---- | ------ |
{{- range $addr := .Addresses }}
| {{ $addr.Addr }} | {{ $addr.AddrType }} | {{ $addr.Vendor }} |
{{- end }}

## Runnig Services
{{- if .Ports }}
| Port | State | Service |
| ---- | ----- | ------- |
{{- range $port := .Ports }}
| {{ $port.Protocol }}/{{ $port.PortID }} | {{ $port.State.State }} | {{ $port.Service }} |
{{- end }}
{{- else }}
Nmap could not find any open port in this host.
{{- end }}

{{- if .ExtraPorts }}

## Extra Ports
| State | Count |
| ----- | ----- |
{{- range $port := .ExtraPorts }}
| {{ $port.State }} | {{ $port.Count }} |
{{- end }}
{{- end }}

## Operating System (OS)
{{- if .Os.OsMatches }}
| Name  | Accuracy |
| ----- | -------- |
{{- range $match := .Os.OsMatches }}
| {{ $match.Name }} | {{ $match.Accuracy }}% |
{{- end }}
{{- else }}
Nmap could not guess the OS of this host.
{{- end }}

## Trace
| Address | Hostname | TTL | RTT |
| ------- | -------- | --- | --- |
{{- range $hop := .Trace.Hops }}
{{- $host := or $hop.Host "unknown" }}
| {{ $hop.IPAddr }} | {{ $host }} | {{ $hop.TTL }} | {{ $hop.RTT }} |
{{- end }}

{{- if .HostScripts }}
## Scripts
{{- range $out := .HostScripts }}
### {{ $out.ID }}
{{- $out.Output }}
{{- range $table := $out.Tables }}
* **{{ $table.Key }}** = {{ $table.Elements }}
{{- end }}
{{- end }}
{{- end }}

{{- if or .IPIDSequence .TCPSequence }}

## Extra Info
{{- if .IPIDSequence }}
IP ID Sequence Generation: {{ .IPIDSequence.Class }}
{{- end }}
{{- if .TCPSequence }}
TCP Sequence Prediction: Difficulty={{ .TCPSequence.Index }} ({{ .TCPSequence.Difficulty }})
{{- end }}
{{- end }}

{{- if ne .Uptime.Seconds 0 }}

## Uptime
{{ .Uptime.Lastboot }} ({{ .Uptime.Seconds }} sec.)
{{- end -}}
