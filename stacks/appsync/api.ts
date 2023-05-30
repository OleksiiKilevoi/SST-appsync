// import { Duration, Expiration } from 'aws-cdk-lib';
// import * as appsync from 'aws-cdk-lib/aws-appsync';
// import * as acm from 'aws-cdk-lib/aws-certificatemanager';
// import * as route53 from 'aws-cdk-lib/aws-route53';
// import * as targets from 'aws-cdk-lib/aws-route53-targets';
import { AppSyncApi, Function, StackContext } from 'sst/constructs';

export function api({ stack, app }: StackContext) {
	// const authorizerFunction = new Function(stack, 'AuthorizerFunction', {
	// 	handler: 'functions/auth/authorizer/main.go',
	// });

	const api = new AppSyncApi(stack, 'AppSyncApi', {
		schema: 'graphql/schema.graphql',
		cdk: {
			graphqlApi: {
				name: `${app.stage}-${app.name}-api`,
				// authorizationConfig: {
				// 	defaultAuthorization: {
				// 		authorizationType: appsync.AuthorizationType.LAMBDA,
				// 		lambdaAuthorizerConfig: {
				// 			handler: authorizerFunction,
				// 			validationRegex: '^ey.*$',
				// 			resultsCacheTtl: app.local ? Duration.seconds(0) : Duration.minutes(5),
				// 		},
				// 	},
				// 	additionalAuthorizationModes: [{
				// 		authorizationType: appsync.AuthorizationType.API_KEY,
				// 		apiKeyConfig: {
				// 			expires: Expiration.after(Duration.days(365)),
				// 		},
				// 	}],
				// },
			},
		},
		resolvers: {
			'Query users': 'functions/users/main.go',
		},
	});

	// if (['dev', 'staging', 'prod'].includes(app.stage)) {
	// 	const apiDomain = `${app.stage}-api.science-card.net`;

	// 	const hostedZone = route53.HostedZone.fromLookup(stack, 'AppSyncApiDomainHostedZone', {
	// 		domainName: apiDomain,
	// 	});

	// 	const cert = new acm.DnsValidatedCertificate(stack, 'AppSyncApiDomainCertificate', {
	// 		hostedZone,
	// 		domainName: apiDomain,
	// 		region: 'us-east-1',
	// 	});

	// 	const appsyncDomainName = new appsync.CfnDomainName(stack, 'AppSyncApiDomainName', {
	// 		certificateArn: cert.certificateArn,
	// 		domainName: apiDomain,
	// 	});

	// 	const assoc = new appsync.CfnDomainNameApiAssociation(stack, 'AppSyncApiDomainNameApiAssociation', {
	// 		apiId: api.apiId,
	// 		domainName: apiDomain,
	// 	});

	// 	assoc.addDependency(appsyncDomainName);

		// new route53.ARecord(stack, 'AppSyncApiDomainAliasRecord', {
		// 	zone: hostedZone,
		// 	recordName: apiDomain,
		// 	target: route53.RecordTarget.fromAlias({
		// 		bind() {
		// 			return {
		// 				hostedZoneId: targets.CloudFrontTarget.CLOUDFRONT_ZONE_ID,
		// 				dnsName: appsyncDomainName.attrAppSyncDomainName,
		// 			};
		// 		},
		// 	}),
		// });
	// }
}
