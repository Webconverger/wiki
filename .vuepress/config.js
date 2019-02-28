module.exports = {
  title: 'Webconverger documentation',
  themeConfig: {
    repo: 'webconverger/wiki',
    docsBranch: 'vuepress',
    editLinks: true,
    displayAllHeaders: true,
    lastUpdated: 'Last Updated'
  },
  evergreen: true,
  plugins: [
    require('vuepress-backlinks')
  ]
}
