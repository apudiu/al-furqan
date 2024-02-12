import { A } from '@solidjs/router';
import { navigationLinks } from '~/lib/routes';

const Navbar = () => {
  return (
    <nav class="min-h-60">
      <ul class="flex flex-col gap-8">
        {navigationLinks.map((navLink) => (
          <li>
            <A
              href={navLink.href}
              class="flex gap-2 text-primary-foreground text-md tracking-wider hover:opacity-90 items-end"
            >
              {navLink.icon && <navLink.icon />}

              <span>{navLink.title}</span>
            </A>
          </li>
        ))}
      </ul>
    </nav>
  );
};
export default Navbar;
