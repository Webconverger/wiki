export default ({ router }) => {
  router.addRoutes([
    { path: '/boot/', redirect: '/API/' },
    { path: '/about/', redirect: '/system_integrators/' },
    { path: '/hetty/', redirect: '/servers/' },
    { path: '/homepage/', redirect: '/API/' },
    { path: '/bugs/', redirect: 'https://github.com/Webconverger/Debian-Live-config/issues' }
  ])
}
