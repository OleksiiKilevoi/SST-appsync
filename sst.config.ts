import { SSTConfig } from "sst";
import { api } from './stacks/appsync/api';
import { dynamo } from "./stacks/dynamo";

export default {
  config(_input) {
    return {
      name: "my-sst-app",
      region: "us-east-1",
    };
  },
  stacks(app) {

    app.setDefaultFunctionProps({
			runtime: 'go1.x',
		});

    app
    .stack(api)
    .stack(dynamo)
  }
} satisfies SSTConfig;
