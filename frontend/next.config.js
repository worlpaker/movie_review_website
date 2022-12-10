/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
};

module.exports = nextConfig;
module.exports = {
  async redirects() {
    return [
      {
        source: '/reviews_profile',
        destination: '/',
        permanent: true,
      },
      {
        source: '/watchlist',
        destination: '/',
        permanent: true,
      },
      {
        source: '/api/:path*',
        destination: 'http://localhost:80/api/:path*', // Redirect to Backend
        permanent: true,
      },
    ];
  },
};