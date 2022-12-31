package sequelize

var packageJson string = `    "rollback": "npx sequelize-cli db:migrate:undo:all",
    "migrate": "npx sequelize-cli db:migrate",
    "new-migration": "npx sequelize-cli migration:create --name=$MIGRATION_NAME",
    "new-seed": "npx sequelize-cli seed:create --name=$SEED_NAME",
    "seed": "npx sequelize-cli db:seed:all",
    "seed-rollback": "npx sequelize-cli db:seed:undo:all"`

var sequelizerc string = `
const path = require('path');

module.exports = {
  'config': path.resolve('config', 'database.js'),
  'models-path': path.resolve('src', 'models'),
  'seeders-path': path.resolve('db', 'seeders'),
  'migrations-path': path.resolve('db', 'migrations')
};
`
