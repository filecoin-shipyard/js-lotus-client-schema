const ndjson = require('ndjson')
const fs = require('fs')
const path = require('path')

function onSchema (schema) {
  const methods = Object.entries(schema).map(([k, v]) => {
    const name = k === 'ID' ? 'id' : k[0].toLowerCase() + k.slice(1)
    const params = (v.params || []).map(p => p.name + ': ' + p.type)
    return `${name} (${params.join(', ')}): ${v.returns}`
  })
  const tmpl = fs.readFileSync(path.join(__dirname, '/LotusRPC.template.d.ts'), 'utf8')
  console.log(tmpl.replace('/* methods */', methods.join('\n  ')))
}

process.stdin.pipe(ndjson.parse()).on('data', onSchema)
