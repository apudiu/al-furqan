import { useLocation } from '@solidjs/router';
import { routePaths } from '~/lib/routes';

const Player = () => {
  const location = useLocation();

  return (
    <section
      class="h-20 fixed left-0 right-0 bottom-0 top-auto bg-primary text-primary-foreground"
      classList={{
        hidden: location.pathname == routePaths.playlist,
      }}
    >
      Player
    </section>
  );
};
export default Player;
