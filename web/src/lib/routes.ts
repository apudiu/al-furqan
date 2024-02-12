import { HomeIcon, LibraryIcon, ListMusicIcon, SearchIcon } from 'lucide-solid';
import { NavigationLink } from '~/types';

export const routePaths = {
  home: '/',
  library: '/library',
  contact: '/contact',
  playlist: '/playlist',
  search: '/search',
  searchQuery: '/search/:*',
};

export const navigationLinks: NavigationLink[] = [
  {
    title: 'Home',
    href: routePaths.home,
    icon: HomeIcon,
  },
  {
    title: 'Search',
    href: routePaths.search,
    icon: SearchIcon,
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

export const searchBarVisiblePaths = [routePaths.search, routePaths.searchQuery];
