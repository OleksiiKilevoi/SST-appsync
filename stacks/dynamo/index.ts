import { StackContext, Table } from 'sst/constructs';

export function dynamo({ stack, app }: StackContext) {
	const users = new Table(stack, 'AccountsTable', {
		cdk: {
			table: {
				tableName: 'users',
			},
		},
		fields: {
			id: 'string',
			email: 'string',
			phoneNumber: 'string',
		},
		primaryIndex: {
			partitionKey: 'id',
		},
		globalIndexes: {
			email: {
				partitionKey: 'email',
			},
			phoneNumber: {
				partitionKey: 'phoneNumber',
			},
		},
	});
	
	app.addDefaultFunctionBinding([users]);

	return {
		usersTable: users,
	};
}
