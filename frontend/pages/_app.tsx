import type { AppProps } from 'next/app';
import Mynavbar from './navbar';
import Authorization from '../components/Authorization';
import '../styles/navbar.css';
import '../styles/home.css';
import '../styles/profile.css';
import '../styles/notification.css';


/*
<Authorization>
</Authorization>

*/

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <Authorization>
        <Mynavbar>
          <Component {...pageProps} />
        </Mynavbar>
      </Authorization>

    </>
  );
}

export default MyApp;
