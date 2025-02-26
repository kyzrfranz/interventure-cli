<main>
    {{ if . }}
      {{ .Bio }}
      <h1>
        {{ .Bio.AcademicTitle }}
        {{ .Bio.FirstName }}
        {{ .Bio.NobilityTitle }}
        {{ .Bio.LastName }}
      </h1>
      <h3>
        {{ .Bio.Elected }}
        {{ .Bio.Party }}
        {{ .Bio.State }}
      </h3>
      {{ if .Bio.Constituency }}
      <h4>
        Wahlkreis: <a href="{{ .Bio.Constituency.URL }}">{{ .Bio.Constituency.Name }}</a>
      </h4>
      {{ end }}
      <img src="{{ .Media.Photo.URL }}" alt="{{ .Media.Photo.AltText }}" title="{{ .Media.Photo.Copyright }}">
      {{ .Bio.BiographicInfo }}

      <h3>Ausschüsse</h3>
      <ul>
        {{ range $name, $cGroup := .Bio.Memberships }}
        <li>
          <h4>{{ $name }}</h4>
          <ol>
            {{ range $cGroup }}
            <li>
              <a href="{{ .URL }}">
                {{ .Name }}
              </a>
            </li>
            {{ end }}
          </ol>
        </li>
        {{ end }}
      </ul>

      <h3>Bezüge</h3>
      {{ .Bio.MandatedPublishableInfo  }}
    {{ else }}
      <h1>Person not found</h1>
    {{ end }}
</main>
