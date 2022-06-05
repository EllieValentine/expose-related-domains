# Console utility that helps to find all related (sub) domains.

Search relays on <a href="https://en.wikipedia.org/wiki/Certificate_Transparency">Certificate Transparency (CT)</a> data.

Queries go to <a href="https://crt.sh/">Crt.sh</a>. The site where you can find all the SSL/TLS certificates of the particular targeted domain.

## How to use

1. Run main.go
2. `Type domain: `<b>github.com</b>

`Example of the response:`

<pre>
examregistration.github.com
vpn-ca.iad.github.com
review-lab.github.com
github.com
www.github.com
pkg.github.com
visualstudio.github.com
www.visualstudio.github.com
render-lab.github.com
www.render-lab.github.com
github.io
docs-front-door.github.com
docs.github.com
smtp.github.com
skyline.github.com
www.skyline.github.com
api.security.github.com
www.api.security.github.com
f.cloud.github.com
registry.github.com
help.github.com
support.enterprise.github.com
www.support.enterprise.github.com
import2.github.com
importer2.github.com
porter2.github.com
jira.github.com
www.jira.github.com
api.stars.github.com
www.api.stars.github.com
fast.github.com
r2.shared.global.fastly.net
mac-installer.github.com
vscode-auth.github.com
classroom.github.com
styleguide.github.com
render.github.com
codespaces-ppe.github.com
codespaces.github.com
codespaces-dev.github.com
lab.github.com
...
</pre>
