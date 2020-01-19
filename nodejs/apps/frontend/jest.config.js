module.exports = {
  name: 'frontend',
  preset: '../../jest.config.js',
  coverageDirectory: '../../coverage/apps/frontend',
  snapshotSerializers: [
    'jest-preset-angular/AngularSnapshotSerializer.js',
    'jest-preset-angular/HTMLCommentSerializer.js'
  ]
};
