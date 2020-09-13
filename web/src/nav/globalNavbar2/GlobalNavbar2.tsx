import * as H from 'history'
import React, { useEffect, useMemo } from 'react'
import { ActivationProps } from '../../../../shared/src/components/activation/Activation'
import { ExtensionsControllerProps } from '../../../../shared/src/extensions/controller'
import { PlatformContextProps } from '../../../../shared/src/platform/context'
import { SettingsCascadeProps } from '../../../../shared/src/settings/settings'
import { AuthenticatedUser } from '../../auth'
import {
    parseSearchURLQuery,
    PatternTypeProps,
    InteractiveSearchProps,
    CaseSensitivityProps,
    CopyQueryButtonProps,
    OnboardingTourProps,
} from '../../search'
import { SearchNavbarItem } from '../../search/input/SearchNavbarItem'
import { showDotComMarketing } from '../../util/features'
import { NavLinks } from '../NavLinks'
import { ThemeProps } from '../../../../shared/src/theme'
import { ThemePreferenceProps } from '../../theme'
import { KeyboardShortcutsProps } from '../../keyboardShortcuts/keyboardShortcuts'
import { QueryState } from '../../search/helpers'
import { InteractiveModeInput } from '../../search/input/interactive/InteractiveModeInput'
import { FiltersToTypeAndValue } from '../../../../shared/src/search/interactive/util'
import { SearchModeToggle } from '../../search/input/interactive/SearchModeToggle'
import { Link } from '../../../../shared/src/components/Link'
import { convertPlainTextToInteractiveQuery } from '../../search/input/helpers'
import { VersionContextDropdown } from '../VersionContextDropdown'
import { VersionContextProps } from '../../../../shared/src/search/util'
import { VersionContext } from '../../schema/site.schema'
import { TelemetryProps } from '../../../../shared/src/telemetry/telemetryService'
import { BrandLogo } from '../../components/branding/BrandLogo'
import { LinkOrSpan } from '../../../../shared/src/components/LinkOrSpan'
import { GlobalMenuPopoverButton } from './menu/GlobalMenuPopoverButton'

interface Props
    extends SettingsCascadeProps,
        PlatformContextProps,
        ExtensionsControllerProps,
        KeyboardShortcutsProps,
        TelemetryProps,
        ThemeProps,
        ThemePreferenceProps,
        ActivationProps,
        PatternTypeProps,
        CaseSensitivityProps,
        InteractiveSearchProps,
        CopyQueryButtonProps,
        VersionContextProps,
        OnboardingTourProps {
    history: H.History
    location: H.Location<{ query: string }>
    authenticatedUser: AuthenticatedUser | null
    authRequired: boolean
    navbarSearchQueryState: QueryState
    onNavbarQueryChange: (queryState: QueryState) => void
    isSourcegraphDotCom: boolean
    isSearchRelatedPage: boolean
    showCampaigns: boolean

    // Whether globbing is enabled for filters.
    globbing: boolean

    /**
     * Which variation of the global navbar to render.
     *
     * 'low-profile' renders the the navbar with no border or background. Used on the search
     * homepage.
     *
     * 'low-profile-with-logo' renders the low-profile navbar but with the homepage logo. Used on repogroup pages.
     */
    variant: 'default' | 'low-profile' | 'low-profile-with-logo' | 'no-search-input'

    splitSearchModes: boolean
    interactiveSearchMode: boolean
    toggleSearchMode: (event: React.MouseEvent<HTMLAnchorElement>) => void
    setVersionContext: (versionContext: string | undefined) => void
    availableVersionContexts: VersionContext[] | undefined

    branding?: typeof window.context.branding

    /** For testing only. Used because reactstrap's Popover is incompatible with react-test-renderer. */
    hideNavLinks: boolean

    setShowGlobalMenu: (value: boolean) => void
}

export const GlobalNavbar2: React.FunctionComponent<Props> = ({
    authRequired,
    isSearchRelatedPage,
    splitSearchModes,
    interactiveSearchMode,
    navbarSearchQueryState,
    versionContext,
    setVersionContext,
    availableVersionContexts,
    caseSensitive,
    patternType,
    onNavbarQueryChange,
    onFiltersInQueryChange,
    hideNavLinks,
    variant,
    isLightTheme,
    branding = window.context?.branding,
    location,
    history,
    setShowGlobalMenu,
    ...props
}) => {
    const query = useMemo(() => parseSearchURLQuery(location.search || ''), [location.search])

    useEffect(() => {
        // In interactive search mode, the InteractiveModeInput component will handle updating the inputs.
        if (!interactiveSearchMode) {
            if (query) {
                onNavbarQueryChange({ query, cursorPosition: query.length })
            } else {
                // If we have no component state, then we may have gotten unmounted during a route change.
                const query = location.state ? location.state.query : ''
                onNavbarQueryChange({
                    query,
                    cursorPosition: query.length,
                })
            }
        }

        if (query) {
            if (!isSearchRelatedPage) {
                // On a non-search related page or non-repo page, we clear the query in
                // the main query input and interactive mode UI to avoid misleading users
                // that the query is relevant in any way on those pages.
                onNavbarQueryChange({ query: '', cursorPosition: 0 })
                onFiltersInQueryChange({})
            }

            if (interactiveSearchMode) {
                let filtersInQuery: FiltersToTypeAndValue = {}
                const { filtersInQuery: newFiltersInQuery, navbarQuery } = convertPlainTextToInteractiveQuery(query)
                filtersInQuery = { ...filtersInQuery, ...newFiltersInQuery }
                onNavbarQueryChange({ query: navbarQuery, cursorPosition: navbarQuery.length })
                onFiltersInQueryChange(filtersInQuery)
            }
        }
    }, [interactiveSearchMode, isSearchRelatedPage, location, onFiltersInQueryChange, onNavbarQueryChange, query])

    return (
        <div className="global-navbar2">
            <GlobalMenuPopoverButton
                logoClassName="global-navbar2__logo"
                isLightTheme={isLightTheme}
                setShowGlobalMenu={setShowGlobalMenu}
            />
            <div className="d-flex w-100" style={{ maxWidth: '600px' }}>
                <SearchNavbarItem
                    {...props}
                    navbarSearchState={navbarSearchQueryState}
                    onChange={onNavbarQueryChange}
                    location={location}
                    history={history}
                    versionContext={versionContext}
                    isLightTheme={isLightTheme}
                    patternType={patternType}
                    caseSensitive={caseSensitive}
                />
            </div>
        </div>
    )
}