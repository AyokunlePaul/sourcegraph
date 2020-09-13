import React from 'react'
import renderer from 'react-test-renderer'
import { setLinkComponent } from '../../../../shared/src/components/Link'
import { ThemePreference } from '../../theme'
import { GlobalNavbar2 } from './GlobalNavbar2'
import { createLocation, createMemoryHistory } from 'history'
import { NOOP_SETTINGS_CASCADE } from '../../../../shared/src/util/searchTestHelpers'
import { SearchPatternType } from '../../graphql-operations'

jest.mock('../search/input/SearchNavbarItem', () => ({ SearchNavbarItem: 'SearchNavbarItem' }))
jest.mock('../components/branding/BrandLogo', () => ({ BrandLogo: 'BrandLogo' }))

const PROPS: React.ComponentProps<typeof GlobalNavbar2> = {
    authenticatedUser: null,
    authRequired: false,
    extensionsController: {} as any,
    location: createLocation('/'),
    history: createMemoryHistory(),
    keyboardShortcuts: [],
    isSourcegraphDotCom: false,
    navbarSearchQueryState: { query: 'q', cursorPosition: 0 },
    onNavbarQueryChange: () => undefined,
    onThemePreferenceChange: () => undefined,
    isLightTheme: true,
    themePreference: ThemePreference.Light,
    patternType: SearchPatternType.literal,
    setPatternType: () => undefined,
    caseSensitive: false,
    setCaseSensitivity: () => undefined,
    platformContext: {} as any,
    settingsCascade: NOOP_SETTINGS_CASCADE,
    showCampaigns: false,
    telemetryService: {} as any,
    hideNavLinks: true, // used because reactstrap Popover is incompatible with react-test-renderer
    filtersInQuery: {} as any,
    splitSearchModes: false,
    interactiveSearchMode: false,
    toggleSearchMode: () => undefined,
    onFiltersInQueryChange: () => undefined,
    isSearchRelatedPage: true,
    copyQueryButton: false,
    versionContext: undefined,
    setVersionContext: () => undefined,
    availableVersionContexts: [],
    variant: 'default',
    globbing: false,
    showOnboardingTour: false,
    branding: undefined,
}

describe('GlobalNavbar2', () => {
    setLinkComponent(props => <a {...props} />)
    afterAll(() => setLinkComponent(() => null)) // reset global env for other tests

    test('default', () => expect(renderer.create(<GlobalNavbar2 {...PROPS} />).toJSON()).toMatchSnapshot())

    test('low-profile', () =>
        expect(renderer.create(<GlobalNavbar2 {...PROPS} variant="low-profile" />).toJSON()).toMatchSnapshot())

    test('no-search-input', () =>
        expect(renderer.create(<GlobalNavbar2 {...PROPS} variant="no-search-input" />).toJSON()).toMatchSnapshot())
})