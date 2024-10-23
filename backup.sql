PGDMP     /    %            	    |            tajalevegsuppliers    15.8    15.8 '               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16388    tajalevegsuppliers    DATABASE     z   CREATE DATABASE tajalevegsuppliers WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C.UTF-8';
 "   DROP DATABASE tajalevegsuppliers;
                postgres    false            �            1259    16459    balance_sheet    TABLE     W  CREATE TABLE public.balance_sheet (
    id integer NOT NULL,
    account_type character varying(50) NOT NULL,
    account_name character varying(255) NOT NULL,
    debit numeric(10,2),
    credit numeric(10,2),
    balance numeric(10,2) NOT NULL,
    transaction_date date NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);
 !   DROP TABLE public.balance_sheet;
       public         heap    postgres    false            �            1259    16458    balance_sheet_id_seq    SEQUENCE     �   CREATE SEQUENCE public.balance_sheet_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.balance_sheet_id_seq;
       public          postgres    false    221                       0    0    balance_sheet_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.balance_sheet_id_seq OWNED BY public.balance_sheet.id;
          public          postgres    false    220            �            1259    16425    bill_details    TABLE       CREATE TABLE public.bill_details (
    id integer NOT NULL,
    bill_number character varying(255) NOT NULL,
    bill_date date NOT NULL,
    bill_total_amount character varying(255) NOT NULL,
    seller_name character varying(255) NOT NULL,
    seller_pan_num character varying(255) NOT NULL,
    customer_name character varying(255) NOT NULL,
    customer_location character varying(255) NOT NULL,
    customer_phone_number character varying(20) NOT NULL,
    customer_pan_container character varying(255) NOT NULL,
    items jsonb
);
     DROP TABLE public.bill_details;
       public         heap    postgres    false                       0    0    TABLE bill_details    ACL     1   GRANT ALL ON TABLE public.bill_details TO romit;
          public          postgres    false    215            �            1259    16424    bill_details_id_seq    SEQUENCE     �   CREATE SEQUENCE public.bill_details_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.bill_details_id_seq;
       public          postgres    false    215                       0    0    bill_details_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.bill_details_id_seq OWNED BY public.bill_details.id;
          public          postgres    false    214                       0    0    SEQUENCE bill_details_id_seq    ACL     ;   GRANT ALL ON SEQUENCE public.bill_details_id_seq TO romit;
          public          postgres    false    214            �            1259    16434    dailyvegetablesales    TABLE     \  CREATE TABLE public.dailyvegetablesales (
    id integer NOT NULL,
    vegetable_name character varying(255) NOT NULL,
    sale_date date NOT NULL,
    quantity_sold numeric(10,2) DEFAULT 0.00,
    rate numeric(10,2) DEFAULT 0.00,
    total_amount numeric(10,2) DEFAULT 0.00,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
 '   DROP TABLE public.dailyvegetablesales;
       public         heap    postgres    false                       0    0    TABLE dailyvegetablesales    ACL     8   GRANT ALL ON TABLE public.dailyvegetablesales TO romit;
          public          postgres    false    217            �            1259    16433    dailyvegetablesales_id_seq    SEQUENCE     �   CREATE SEQUENCE public.dailyvegetablesales_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 1   DROP SEQUENCE public.dailyvegetablesales_id_seq;
       public          postgres    false    217                       0    0    dailyvegetablesales_id_seq    SEQUENCE OWNED BY     Y   ALTER SEQUENCE public.dailyvegetablesales_id_seq OWNED BY public.dailyvegetablesales.id;
          public          postgres    false    216                       0    0 #   SEQUENCE dailyvegetablesales_id_seq    ACL     B   GRANT ALL ON SEQUENCE public.dailyvegetablesales_id_seq TO romit;
          public          postgres    false    216            �            1259    16447    ledger_entries    TABLE     s  CREATE TABLE public.ledger_entries (
    id integer NOT NULL,
    date date NOT NULL,
    account character varying(255) NOT NULL,
    billnumber character varying(255) NOT NULL,
    debit numeric(10,2) DEFAULT 0.00,
    credit numeric(10,2) DEFAULT 0.00,
    balance_amount numeric(10,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
 "   DROP TABLE public.ledger_entries;
       public         heap    postgres    false                       0    0    TABLE ledger_entries    ACL     3   GRANT ALL ON TABLE public.ledger_entries TO romit;
          public          postgres    false    219            �            1259    16446    ledger_entries_id_seq    SEQUENCE     �   CREATE SEQUENCE public.ledger_entries_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.ledger_entries_id_seq;
       public          postgres    false    219                       0    0    ledger_entries_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.ledger_entries_id_seq OWNED BY public.ledger_entries.id;
          public          postgres    false    218                       0    0    SEQUENCE ledger_entries_id_seq    ACL     =   GRANT ALL ON SEQUENCE public.ledger_entries_id_seq TO romit;
          public          postgres    false    218            l           2604    16462    balance_sheet id    DEFAULT     t   ALTER TABLE ONLY public.balance_sheet ALTER COLUMN id SET DEFAULT nextval('public.balance_sheet_id_seq'::regclass);
 ?   ALTER TABLE public.balance_sheet ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    221    220    221            b           2604    16428    bill_details id    DEFAULT     r   ALTER TABLE ONLY public.bill_details ALTER COLUMN id SET DEFAULT nextval('public.bill_details_id_seq'::regclass);
 >   ALTER TABLE public.bill_details ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    214    215            c           2604    16437    dailyvegetablesales id    DEFAULT     �   ALTER TABLE ONLY public.dailyvegetablesales ALTER COLUMN id SET DEFAULT nextval('public.dailyvegetablesales_id_seq'::regclass);
 E   ALTER TABLE public.dailyvegetablesales ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    216    217            h           2604    16450    ledger_entries id    DEFAULT     v   ALTER TABLE ONLY public.ledger_entries ALTER COLUMN id SET DEFAULT nextval('public.ledger_entries_id_seq'::regclass);
 @   ALTER TABLE public.ledger_entries ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    219    219                      0    16459    balance_sheet 
   TABLE DATA           }   COPY public.balance_sheet (id, account_type, account_name, debit, credit, balance, transaction_date, created_at) FROM stdin;
    public          postgres    false    221   �.                 0    16425    bill_details 
   TABLE DATA           �   COPY public.bill_details (id, bill_number, bill_date, bill_total_amount, seller_name, seller_pan_num, customer_name, customer_location, customer_phone_number, customer_pan_container, items) FROM stdin;
    public          postgres    false    215   �.       	          0    16434    dailyvegetablesales 
   TABLE DATA           {   COPY public.dailyvegetablesales (id, vegetable_name, sale_date, quantity_sold, rate, total_amount, created_at) FROM stdin;
    public          postgres    false    217   �;                 0    16447    ledger_entries 
   TABLE DATA           r   COPY public.ledger_entries (id, date, account, billnumber, debit, credit, balance_amount, created_at) FROM stdin;
    public          postgres    false    219   �B                  0    0    balance_sheet_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.balance_sheet_id_seq', 1, false);
          public          postgres    false    220                       0    0    bill_details_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.bill_details_id_seq', 21, true);
          public          postgres    false    214                        0    0    dailyvegetablesales_id_seq    SEQUENCE SET     K   SELECT pg_catalog.setval('public.dailyvegetablesales_id_seq', 5764, true);
          public          postgres    false    216            !           0    0    ledger_entries_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.ledger_entries_id_seq', 3, true);
          public          postgres    false    218            w           2606    16465     balance_sheet balance_sheet_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.balance_sheet
    ADD CONSTRAINT balance_sheet_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.balance_sheet DROP CONSTRAINT balance_sheet_pkey;
       public            postgres    false    221            o           2606    16432    bill_details bill_details_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.bill_details
    ADD CONSTRAINT bill_details_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.bill_details DROP CONSTRAINT bill_details_pkey;
       public            postgres    false    215            q           2606    16443 ,   dailyvegetablesales dailyvegetablesales_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.dailyvegetablesales
    ADD CONSTRAINT dailyvegetablesales_pkey PRIMARY KEY (id);
 V   ALTER TABLE ONLY public.dailyvegetablesales DROP CONSTRAINT dailyvegetablesales_pkey;
       public            postgres    false    217            s           2606    16445 D   dailyvegetablesales dailyvegetablesales_vegetable_name_sale_date_key 
   CONSTRAINT     �   ALTER TABLE ONLY public.dailyvegetablesales
    ADD CONSTRAINT dailyvegetablesales_vegetable_name_sale_date_key UNIQUE (vegetable_name, sale_date);
 n   ALTER TABLE ONLY public.dailyvegetablesales DROP CONSTRAINT dailyvegetablesales_vegetable_name_sale_date_key;
       public            postgres    false    217    217            u           2606    16457 "   ledger_entries ledger_entries_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.ledger_entries
    ADD CONSTRAINT ledger_entries_pkey PRIMARY KEY (id);
 L   ALTER TABLE ONLY public.ledger_entries DROP CONSTRAINT ledger_entries_pkey;
       public            postgres    false    219                  x������ � �           x��]Ko�>�~ŀ'0�����(%���Hr#�In��.w�]�"`H���0���Ȳ` �r�������.�SU3=;��:H|�U���W_W7�#w�:��E�h&��L�ΦgO�g?M��ӳ����^OϞOϿN�s�fz����U��kgz�bz�lz�]�������?:����C��f�I��uz������96=�u�0��Bɂx ��ϗ���66k�r�n�I?�s��;rߛ�}��~������>p�v�(�����(�)�|<2���3�f��Ao��̾����%w�n4�O��C�j6��p�ed �s�!��2�o��hx��~��=^�W`�";p	�E.��oMv&�٥�&ݒxcNߗ,�&�jQP�O�H�����D�O���LϾS��= \yD6��~ԏ��L��*���x�UV��Z�U˯�� :6����goS󍕜d5x��|��ԆDVŦ�3G�b�$n' ��� ��1Ѭ���s4L��i��[dH`�C�Â�6��q��l��̖���~�M���x?8�9Y.4��vS��W��$��X�3�Э�_���3}��n���t�__�gn2�d{�Ncpu*��N�?����;������q��M��X�����p��>�7��M�����`��� }?	_X�Be����M$4j�^^�� �&f�~dv�<�\�����mR�+)�k���=c�}��i,ݠY�ybY8�GU���� =�3;��b�("��Z�2��r��c�+��C�@���d��`mC��9t�9��f9��7�#s8� ����<r�#�	��B� �<���(`Ԭ�M`H�r,�>6��#,CeЄ��I܀�
M
��9.�UB&@8�>��d�D;�DB���AY��b��qU�piA
�7i�i<v���Ps�I�.�	��W �`��^���!��U��`�L�v�	z��|��}Tl'JȲP�JP+��3;4H4��ښ��9h��OR��6��7���l��Db�~_=�_2�bd��B���k�Z'6��ᓭ�l|��Os����h�Û��x㳍���w66�ol}z����b�zP��F��:SbQ@�����&
^dʧY������� ���䈬F� ��J���ɑ�]@�+������AE/g~}�#�$f���H	¹�f{���>M��%L��o�Q]ٓ�B�xU�� bK��6��hP�4�7��Ab�S�h ��s1
Z
�����VtA"�A�>�A�/#.��y�9 ��3t3�9�P��1��v��h�CQQ-�Dq�Xy:B�P���4Bdy��D��Mi�~	Ë�&W�8CV��!(ah�-':DM��	�y�5LVP<�Ȋ�9���zu����i~���3�����4I��&���p*,1KU?!َ;t�AX~BBУF孛
�
�9��#9��H`��8�]gV�"�@z4wxe�%Q�f,��ٻ�&��6�v�  tˀ
�k��6���,N�6 \c�(�M �z����@f�=�vn��Qw����"���f��!�b����e��E��ª����ޝs��*���~��^vJ������d��tJ�h�C N��E��7@U�7�9F%��&C@HOBa��8<�wg�V�nH[�A�!8���|�]��e�N����#s8�Z�s�֬'�~!BFH"!~�G�p�&<Gh����g�x�/:�MF�fޡC	�u(�Ry��6�D	}�4� !M`m$l��a��N{D�G�W�&+�vݾOY��t7R~K4�7�'�%�.k�$[�c@+� nOhβ4u��X�FiSm�p�R��4�ab=:frҲD3$\�e�5`(G��&{#�doؕ��H������H�-�����[�~wt<7���R���mk�c��P۹DNK���7��7�89Y�І�I�%��G���w�ێ���)>�
�6ʽ9��gL��ؘ��)͐pu��sI���q�ͤm�q�����
g�`G1�X�$`<�&g����a/J��) �+@7xؙ���B|Q��R�);ؔl�6��2�(q���
M4�A��W��p�\a�:����z��Z߯��n���e2�H6�'���ޟ�����V0�`��h�p(��P,�ݒ�4��B<E����m+V�0��X�F���>T^��K���ղ�qr4�F�/����q5��ܑ����|F�wAIK�(��p���y��h�%=QH����Z�~��oU��j�w��ӎXD����~K�g�8��޸�T���ώG�<�B�m��Ȗ�E����*#�"��Y+�<�6��M@�_�V2��`����J��코+)> Ѕ���j�j�2���Y��UNl��:Ɋl���1θ
t�6�}���I\RF��8}G��,1�� ΆL�dN��	G�Ɍ.��rn�s7\��{�ip�/���0�a����"ڔ3�x���ʠ�8�J�|�B���X��l�t�:��)D+VS-ٓ�km!����u��*=�1�b|8�{�3��ͫT���� �fFJQB�ШХ��]'��)(�Kc�+P"2�M���1 "�a�4`�Z�𑣗�[�r$>
>��[kֳ'\�0�x|�.�h0,F(���|(�bB. ���e���%�Y���o!�5l/�9�5�RA������P���~�`���32�ss�TK��e�.�r�"���R����(R��ҊSPbP�sq-���>�,�m�bǌv�A�cn�k��s� �d�Y'\�5r+���ϯh�UI�*	s�x^I(���%O��1&W�}�E�~��s�]F�*z{����{��O脼A酟�!A��ݾAG��� ���J��HK�}A�}��ީ�j�C.)��4���å��N����w���X������C�/��/
�ToV�����I{>S�A�n�B�*��E��~�yRI�J�^b+E��Vs�}G�Qde�˴8],~̡8i
���(Z��/��K�՟��)4��,"t��9W�ա=4��BG#a�k��zN4�M�x[��w��=f�mw����W�[&����銓�����u��p��=E<&d8�K�l���x�E@��8�4�s����/d��+�.�VmX'hW�LP�K��9q��;$<�[N�̵
��1X��ޣ��Ҳ#�%�t1N�]%qr�-a�_"���4A<���{��(vsSd)��gzP	_xl�vBG�>�N��rh�*RWE*�z��"�O�7n�jY{F      	   �  x��XKkW^_��Y&���~d׸�Bj�.���U{�F����@���5�@)�i�#	�B�q��~JϹw�yH3nV��o����s�(e4y�N�GErg��q�|��>_-�Y-o�*�)����QCJ	�?����uB�}��+5tZ;P�*�?�3D�\]?[-��j���:(h��>vbZx�(-(�_����H_��W��o�@R�����C��j���� �B-���~�H�]Hq�Op�!V*Ȟ?8��#�/���y��ju}�Z�������!�e�1sf,<�HF�>���ב�	���3f�_w#9a��X��t��X?�@��ꗣr�ʄ�����(9��m�!�Q� �!W�)eO{��z�ȉ/ҳ<��l1�bѦ��P�3i1,OO�	f�'R��n�!Ș�)��3S%�v���lZ���㠩F��&jlW��C&�IA ȼ�,�Z�	_�3���J b����E��"�?Q���0��F2NN}1�|����H�C7�t
���ʏ�y�G_x�Z.�,�=�m�Y*S��e�-$��$��G��<�Mh>��ho�(&��)��)��c��Q���s�|�x��/a*��g;�H����(�@
��-�ӨeZg��2 ױi`o�(���8�h���|}���?bZ���$�	����.�*t|����o��I'��B�8-lgDiС$J�n��Z*�Qd�F燢���� hl��a! ��^U�r�S�Ȇ�e�{��t�2ӐOO�$=��=̦�k;=����Ѷ��(EZ'��D���yX�F�d�$Ͳt4���_�|\{>�u��c�]�.�β�V14	9�5�\3�����h�B#%�����O��*Ѐ�n�{*�wX
�[�[�������(�������������[�����p��"�����s�5��M�˻��bh��ЅΑ��!Bq����c_d���˭�s�r�C�jKj�J����[+;��g�t���$F�(`�=���PH��Ϛ���Āڣ*'�
F��"�T5�2dt�h��������SV���|��I�}X`˵#i��Uk�d}���kG��yɦ���6��ˇ�B1z(�P-���&���0�Ҹ���jr��Յq��$\�xr�����Q�i�0��Ӱ����c�o]��B�2M��8o��jZ��zܜ2��aW�Sk-�e�$�Q*A�ʂ�A�t��_7gp�c�z�w�+*���fݒ;j�T;���:��O� r{�lT��xV��p����n�L��'%l�8���ϰ
�PZ������2�0�` 8�t�n�pV��ߴl��RX>J��qa��\y<�Y�T��I�q��3�6C6���	qH<�"F�9>�E��v\���4�p��EA�y��Zq���:���.,�*s�{Sx�=A����,�s �{�[qCI�maÒ�dx�`a�;Jg'3��i��[uv��V��ɵŅ�C}Q��XL��V�h�-E�ӽ�p)d�Yr|��.�&��^Ȟ��))U���Y�6�TQ}m�B�`1�4?��t�u%p+j	�}m���ނ22�)�UM�x�m^���054�]s�@U��FgQ5��c�&llϚ
�a������l��H2Y��G��o���k��6�Wt�dX)�r���%�&���X]�n�nC�ɀf�������V����[V��p�:����u��            x������ � �     