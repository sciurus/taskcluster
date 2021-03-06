const path = require('path');
const { Schema } = require('taskcluster-lib-postgres');
const {readRepoFile, writeRepoFile} = require('../../utils');

exports.tasks = [{
  title: 'README Stored Procedures',
  requires: ['db-schema-serializable'],
  provides: ['readme-stored-procedures'],
  run: async (requirements, utils) => {
    const schema = Schema.fromSerializable(requirements['db-schema-serializable']);
    const methods = schema.allMethods().filter(method => !method.deprecated);
    const serviceNames = [...new Set([...methods].map(({ serviceName }) => serviceName).sort())];
    const services = new Map();

    serviceNames.forEach(sn => {
      const serviceMethods = [...methods].reduce((acc, method) => {
        if (method.serviceName !== sn) {
          return acc;
        }

        return acc.concat(method);
      }, []);

      services.set(sn, serviceMethods.sort((a, b) => a.name.localeCompare(b.name)));
    });

    const sections = [...services.entries()].map(([serviceName, methods]) => {
      return [
        `## ${serviceName}`,
        '',
        '| Name | Mode | Arguments | Returns | Description |',
        '| --- | --- | --- | --- | --- |',
        ...[...methods.map(method => `| ${method.name} | ${method.mode} | ${method.args} | ${method.returns} | ${method.description.replace(/\n/g, '<br />')} |`)],
      ].join('\n');
    });

    const content = await readRepoFile(path.join('db', 'fns.md'));
    const newContent = content.replace(
      /(<!-- SP BEGIN -->)(?:.|\n)*(<!-- SP END -->)/m,
      `$1\n${sections.join('\n')}\n$2`);

    if (content !== newContent) {
      await writeRepoFile(path.join('db', 'fns.md'), newContent);
    }
  },
}];
