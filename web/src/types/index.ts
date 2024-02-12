import { LucideProps } from 'lucide-solid/dist/types/types';
import { JSX } from 'solid-js';

export type NavigationLink = {
  title: string;
  href: string;
  icon?: (props: LucideProps) => JSX.Element;
};
