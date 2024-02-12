import { HomeIcon, LibraryIcon, ListMusicIcon } from 'lucide-solid';
import { NavigationLink } from '~/types';

export const routePaths = {
  home: '/',
  library: '/library',
  contact: '/contact',
  playlist: '/playlist',
};

export const navigationLinks: NavigationLink[] = [
  {
    title: 'Home',
    href: routePaths.home,
    icon: HomeIcon,
  },
  {
    title: 'Library',
    href: routePaths.library,
    icon: LibraryIcon,
  },
  {
    title: 'Playlist',
    href: routePaths.playlist,
    icon: ListMusicIcon,
  },
];
