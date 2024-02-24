import { FolderIcon, HomeIcon, UsersIcon } from '@heroicons/react/24/outline';
import { useState } from 'react';

const navigation = [
  { name: 'Dashboard', href: '#', icon: HomeIcon },
  { name: 'Users', href: '#', icon: UsersIcon },
  { name: 'Products', href: '#', icon: FolderIcon }
];

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(' ');
}

const SideBar = () => {
  const [currentNav, setCurrentNav] = useState<String>('Dashboard');


  return (
    <div className="flex flex-col h-full gap-y-5 w-64 w-max-64 overflow-auto bg-indigo-600 px-6">
      <div className="flex h-16 shrink-0 items-center">
        <img
          className="h-8 w-auto"
          src="https://tailwindui.com/img/logos/mark.svg?color=white"
          alt="Your Company"
        />
      </div>

      <nav className="flex flex-1 flex-col">
        <ul role="list" className="flex flex-1 flex-col gap-y-7">
          <li>
            <ul role="list" className="-mx-2 space-y-1">

              {navigation.map(item => (
                <li key={item.name}>
                  <a
                    href={item.href}
                    onClick={() => {setCurrentNav(item.name)}}
                    className={classNames(
                      item.name === currentNav
                        ? 'bg-indigo-700 text-white'
                        : 'text-indigo-200 hover:text-white hover:bg-indigo-700',
                      'group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'
                    )}
                  >
                    <item.icon
                      className={'h-6 w-6 shrink-0'}
                      aria-hidden="true"
                    />
                    {item.name}
                  </a>
                </li>

              ))}
            </ul>
          </li>
        </ul>
      </nav>
    </div>
  );
};

export default SideBar;
