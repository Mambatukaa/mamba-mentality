import { FolderIcon, HomeIcon, UsersIcon } from '@heroicons/react/24/outline';
import { Link } from 'react-router-dom';
import { useState } from 'react';

const navigation = [
  { name: 'Home', href: '/home', icon: HomeIcon },
  { name: 'Users', href: '/users', icon: UsersIcon },
  { name: 'Products', href: '/products', icon: FolderIcon }
];

const SideBar = () => {
  const [currentNav, setCurrentNav] = useState<String>('Home');

  return (
    <nav className="flex flex-col bg-gallery-100 p-2">
      <ul role="list" className="space-y-1">
        {navigation.map(item => (
          <li key={item.name}>
            <Link
              to={item.href}
              onClick={() => {
                setCurrentNav(item.name);
              }}
              className={`menu-item ${
                item.name === currentNav ? 'bg-white' : ''
              }`}
            >
              <item.icon className={'h-3 w-3 shrink-0'} aria-hidden="true" />
              {item.name}
            </Link>
          </li>
        ))}
      </ul>
    </nav>
  );
};

export default SideBar;
