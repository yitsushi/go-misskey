import jinja2
import glob
import re
import os

template = jinja2.Template("""package {{directory}}_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_{{function}}() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Drive().{{module}}().{{function}}()
	if err != nil {
		log.Printf("[Drive/{{module}}/{{function}}] %s", err)

		return
	}

	log.Printf("[Drive/{{module}}/{{function}}] %s ____", resp.Name)
}
""")

for f in glob.glob("services/drive/*/*.go"):
    directory, part = f.split('/')[-2:]

    part = part[:-3]
    module = directory[:-1].capitalize()
    targetFile = f'services/drive/{directory}/{part}_test.go'

    if os.path.isfile(targetFile):
        continue

    try:
        with open(f, "r") as s:
            function = re.search(r"func \(s \*Service\) ([A-Za-z]+)", s.read()).group(1)
    except:
        continue

    with open(targetFile, "w") as t:
        t.write(template.render(directory=directory, module=module, function=function))
