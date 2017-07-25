# examples/scenarios/gh

Serverless GitHub webhooks FTW!

#### Configuration

The following five config settings should be configured prior to running this example:

```
lumi config platform:config:region EastUS
lumi config github:config:repo     pulumi/lumi
lumi config github:config:token    <token from https://github.com/settings/tokens>
lumi config github:config:user     lukehoban
lumi config gh:config:slackToken   <token from https://api.slack.com/custom-integrations/legacy-tokens>
```